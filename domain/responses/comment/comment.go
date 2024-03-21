package comment

import (
	"mygram-api/domain/responses/photo"
	"mygram-api/domain/responses/user"
)

type CommentResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
}

type CommentWithRelationResponse struct {
	ID      uint                       `json:"id"`
	Message string                     `json:"message"`
	PhotoID uint                       `json:"photo_id"`
	UserID  uint                       `json:"user_id"`
	User    user.UserRelationsResponse `json:"user"`
	Photo   photo.PhotoResponse        `json:"photo"`
}
