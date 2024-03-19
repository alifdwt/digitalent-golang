package service

import (
	"errors"
	"mygram-api/domain/requests/auth"
	"mygram-api/domain/responses"
	userRes "mygram-api/domain/responses/user"
	"mygram-api/mapper"
	"mygram-api/pkg/dotenv"
	"mygram-api/pkg/hashing"
	"mygram-api/pkg/logger"
	"mygram-api/pkg/token"
	"mygram-api/repository"
)

type authService struct {
	config dotenv.Config
	user   repository.UserRepository
	hash   hashing.Hashing
	log    logger.Logger
	token  token.Maker
	mapper mapper.UserMapping
}

func NewAuthService(config dotenv.Config, user repository.UserRepository, hash hashing.Hashing, log logger.Logger, token token.Maker, mapper mapper.UserMapping) *authService {
	return &authService{
		config: config,
		user:   user,
		hash:   hash,
		log:    log,
		token:  token,
		mapper: mapper,
	}
}

func (s *authService) Register(input *auth.RegisterRequest) (*userRes.UserResponse, error) {
	hashing, err := s.hash.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = hashing

	user, err := s.user.CreateUser(input)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(user)

	return mapper, nil
}

func (s *authService) Login(input *auth.LoginRequest) (*responses.Token, error) {
	res, err := s.user.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("error while get user: " + err.Error())
	}

	err = s.hash.ComparePassword(res.Password, input.Password)
	if err != nil {
		return nil, errors.New("error while compare password: " + err.Error())
	}

	accessToken, err := s.createAccessToken(int(res.ID), res.Email, res.Username, int(res.Age))
	if err != nil {
		return nil, err
	}

	return &responses.Token{
		Token: accessToken,
	}, nil
}

func (s *authService) createAccessToken(id int, email string, username string, age int) (string, error) {
	res, err := s.token.CreateToken(id, email, username, age, "access", s.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return "", err
	}

	return res, nil
}
