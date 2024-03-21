package mapper

import (
	"mygram-api/domain/responses/user"
	"mygram-api/models"
)

type userMapper struct {
}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) ToUserResponse(request *models.User) *user.UserResponse {
	return &user.UserResponse{
		ID:              request.ID,
		Email:           request.Email,
		Username:        request.Username,
		Age:             request.Age,
		ProfileImageURL: request.ProfileImageURL,
	}
}

func (m *userMapper) ToUserRelationsResponse(request *models.User) *user.UserRelationsResponse {
	return &user.UserRelationsResponse{
		ID:       request.ID,
		Email:    request.Email,
		Username: request.Username,
	}
}
