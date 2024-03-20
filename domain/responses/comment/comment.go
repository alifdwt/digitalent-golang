package comment

type CommentResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
}
