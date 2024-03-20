package service

import (
	"mygram-api/mapper"
	"mygram-api/pkg/dotenv"
	"mygram-api/pkg/hashing"
	"mygram-api/pkg/logger"
	"mygram-api/pkg/token"
	"mygram-api/repository"
)

type Service struct {
	Auth    AuthService
	User    UserService
	Photo   PhotoService
	Comment CommentService
}

type Deps struct {
	Config     dotenv.Config
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	TokenMaker token.Maker
	Logger     logger.Logger
	Mapper     mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:    NewAuthService(deps.Config, deps.Repository.User, deps.Hashing, deps.Logger, deps.TokenMaker, deps.Mapper.UserMapper),
		User:    NewUserService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Mapper.UserMapper),
		Photo:   NewPhotoService(deps.Repository.Photo, deps.Mapper.PhotoMapper),
		Comment: NewCommentService(deps.Repository.Comment, deps.Mapper.CommentMapper),
	}
}
