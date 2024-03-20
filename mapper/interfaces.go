package mapper

import (
	"mygram-api/domain/responses/comment"
	"mygram-api/domain/responses/photo"
	"mygram-api/domain/responses/user"
	"mygram-api/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *user.UserResponse
}

type PhotoMapping interface {
	ToPhotoResponse(request *models.Photo) *photo.PhotoResponse
	ToPhotoResponses(requests *[]models.Photo) []photo.PhotoResponse
}

type CommentMapping interface {
	ToCommentResponse(request *models.Comment) *comment.CommentResponse
	ToCommentResponses(requests *[]models.Comment) []comment.CommentResponse
}
