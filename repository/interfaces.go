package repository

import (
	"mygram-api/domain/requests/auth"
	"mygram-api/domain/requests/photo"
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

type PhotoRepository interface {
	GetPhotoAll() (*[]models.Photo, error)
	GetPhotoById(photoId uint) (*models.Photo, error)
	CreatePhoto(userId uint, request photo.CreatePhotoRequest) (*models.Photo, error)
	UpdatePhoto(userId uint, photoId uint, updatedPhoto photo.UpdatePhotoRequest) (*models.Photo, error)
	DeletePhoto(photoId uint) (*models.Photo, error)
}
