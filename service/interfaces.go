package service

import (
	"mygram-api/domain/requests/auth"
	"mygram-api/domain/requests/comment"
	"mygram-api/domain/requests/photo"
	socialmedia "mygram-api/domain/requests/social_media"
	"mygram-api/domain/requests/user"
	"mygram-api/domain/responses"
	commentRes "mygram-api/domain/responses/comment"
	photoRes "mygram-api/domain/responses/photo"
	userRes "mygram-api/domain/responses/user"
	"mygram-api/models"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*userRes.UserResponse, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
}

type UserService interface {
	UpdateUserById(id uint, request *user.UpdateUserRequest) (*userRes.UserResponse, error)
	DeleteUserById(id uint) (*userRes.UserResponse, error)
}

type PhotoService interface {
	CreatePhoto(userId uint, request photo.CreatePhotoRequest) (*photoRes.PhotoResponse, error)
	GetPhotoAll() (*[]photoRes.PhotoResponse, error)
	GetPhotoById(photoId uint) (*photoRes.PhotoResponse, error)
	UpdatePhoto(userId uint, photoId uint, request photo.UpdatePhotoRequest) (*photoRes.PhotoResponse, error)
	DeletePhoto(photoId uint) (*photoRes.PhotoResponse, error)
}

type CommentService interface {
	CreateComment(userId uint, request comment.CreateCommentRequest) (*commentRes.CommentResponse, error)
	GetCommentAll() (*[]commentRes.CommentResponse, error)
	GetCommentById(commentId uint) (*commentRes.CommentResponse, error)
	UpdateComment(userId uint, commentId uint, request comment.UpdateCommentRequest) (*commentRes.CommentResponse, error)
	DeleteComment(commentId uint) (*commentRes.CommentResponse, error)
}

type SocialMediaService interface {
	CreateSocialMedia(userId uint, request socialmedia.CreateSocialMediaRequest) (*models.SocialMedia, error)
	GetSocialMediaAll(userId uint) (*[]models.SocialMedia, error)
	GetSocialMediaById(socialMediaId uint) (*models.SocialMedia, error)
	UpdateSocialMedia(socialMediaId uint, request socialmedia.UpdateSocialMediaRequest) (*models.SocialMedia, error)
	DeleteSocialMedia(socialMediaId uint) (*models.SocialMedia, error)
}
