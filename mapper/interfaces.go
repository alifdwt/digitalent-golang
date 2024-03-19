package mapper

import (
	"mygram-api/domain/responses/user"
	"mygram-api/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *user.UserResponse
}
