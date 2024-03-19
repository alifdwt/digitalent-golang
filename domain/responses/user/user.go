package user

type UserResponse struct {
	ID              uint   `json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Age             uint   `json:"age"`
	ProfileImageURL string `json:"profile_image_url"`
}
