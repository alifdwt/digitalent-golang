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
	socmedRes "mygram-api/domain/responses/social_media"
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

type PhotoService interface {
	CreatePhoto(userId uint, request photo.CreatePhotoRequest) (*photoRes.PhotoResponse, error)
	GetPhotoAll() (*[]photoRes.PhotoWithRelationResponse, error)
	GetPhotoById(photoId uint) (*photoRes.PhotoWithRelationResponse, error)
	UpdatePhoto(userId uint, photoId uint, request photo.UpdatePhotoRequest) (*photoRes.PhotoResponse, error)
	DeletePhoto(photoId uint) (*photoRes.PhotoResponse, error)
}

type CommentService interface {
	CreateComment(userId uint, request comment.CreateCommentRequest) (*commentRes.CommentResponse, error)
	GetCommentAll() (*[]commentRes.CommentWithRelationResponse, error)
	GetCommentById(commentId uint) (*commentRes.CommentWithRelationResponse, error)
	UpdateComment(userId uint, commentId uint, request comment.UpdateCommentRequest) (*commentRes.CommentResponse, error)
	DeleteComment(commentId uint) (*commentRes.CommentResponse, error)
}

type SocialMediaService interface {
	CreateSocialMedia(userId uint, request socialmedia.CreateSocialMediaRequest) (*socmedRes.SocialMediaResponse, error)
	GetSocialMediaAll(userId uint) (*[]socmedRes.SocialMediaWithRelationResponse, error)
	GetSocialMediaById(socialMediaId uint) (*socmedRes.SocialMediaWithRelationResponse, error)
	UpdateSocialMedia(socialMediaId uint, request socialmedia.UpdateSocialMediaRequest) (*socmedRes.SocialMediaResponse, error)
	DeleteSocialMedia(socialMediaId uint) (*socmedRes.SocialMediaResponse, error)
}
