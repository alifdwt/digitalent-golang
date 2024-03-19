package service

import (
	"mygram-api/domain/requests/auth"
	"mygram-api/domain/requests/user"
	"mygram-api/domain/responses"
	userRes "mygram-api/domain/responses/user"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*userRes.UserResponse, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
}

type UserService interface {
	UpdateUserById(id uint, request *user.UpdateUserRequest) (*userRes.UserResponse, error)
	DeleteUserById(id uint) (*userRes.UserResponse, error)
}
