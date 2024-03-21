package photo

import "mygram-api/domain/responses/user"

type PhotoResponse struct {
	ID       uint   `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type PhotoWithRelationResponse struct {
	ID       uint                       `json:"id"`
	Caption  string                     `json:"caption"`
	Title    string                     `json:"title"`
	PhotoURL string                     `json:"photo_url"`
	UserID   uint                       `json:"user_id"`
	User     user.UserRelationsResponse `json:"user"`
}
