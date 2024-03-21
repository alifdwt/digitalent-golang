package mapper

import (
	"mygram-api/domain/responses/comment"
	"mygram-api/domain/responses/photo"
	socialmedia "mygram-api/domain/responses/social_media"
	"mygram-api/domain/responses/user"
	"mygram-api/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *user.UserResponse
	ToUserRelationsResponse(request *models.User) *user.UserRelationsResponse
}

type PhotoMapping interface {
	ToPhoto(request *models.Photo) *photo.PhotoResponse
	ToPhotoResponse(request *models.Photo) *photo.PhotoWithRelationResponse
	ToPhotoResponses(requests *[]models.Photo) []photo.PhotoWithRelationResponse
}

type CommentMapping interface {
	ToCommentResponse(request *models.Comment) *comment.CommentResponse
	ToCommentWithRelationResponse(request *models.Comment) *comment.CommentWithRelationResponse
	ToCommentWithRelationResponses(requests *[]models.Comment) []comment.CommentWithRelationResponse
}

type SocialMediaMapping interface {
	ToSocialMediaResponse(request *models.SocialMedia) *socialmedia.SocialMediaResponse
	ToSocialMediaWithRelationResponse(request *models.SocialMedia) *socialmedia.SocialMediaWithRelationResponse
	ToSocialMediaWithRelationResponses(requests *[]models.SocialMedia) []socialmedia.SocialMediaWithRelationResponse
}
