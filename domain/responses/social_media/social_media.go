package socialmedia

import "mygram-api/domain/responses/user"

type SocialMediaResponse struct {
	ID             uint              `json:"id"`
	Name           string            `json:"name"`
	SocialMediaURL string            `json:"social_media_url"`
	UserID         uint              `json:"user_id"`
	User           user.UserResponse `json:"user"`
}
