package service

import (
	"mygram-api/domain/requests/user"
	userRes "mygram-api/domain/responses/user"
	"mygram-api/mapper"
	"mygram-api/pkg/hashing"
	"mygram-api/pkg/logger"
	"mygram-api/repository"
)

type userService struct {
	Repository repository.UserRepository
	hash       hashing.Hashing
	log        logger.Logger
	mapper     mapper.UserMapping
}

func NewUserService(user repository.UserRepository, hash hashing.Hashing, logger logger.Logger, mapper mapper.UserMapping) *userService {
	return &userService{
		Repository: user,
		hash:       hash,
		log:        logger,
		mapper:     mapper,
	}
}

func (s *userService) UpdateUserById(id uint, request *user.UpdateUserRequest) (*userRes.UserResponse, error) {
	res, err := s.Repository.UpdateUserById(id, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) DeleteUserById(id uint) (*userRes.UserResponse, error) {
	res, err := s.Repository.DeleteUserById(id)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}
