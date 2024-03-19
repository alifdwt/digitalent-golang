package repository

import (
	"mygram-api/domain/requests/auth"
	"mygram-api/domain/requests/user"
	"mygram-api/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	CreateUser(registerReq *auth.RegisterRequest) (*models.User, error)
	UpdateUserById(id uint, updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id uint) (*models.User, error)
}
