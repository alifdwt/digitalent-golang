package mapper

import (
	"mygram-api/domain/responses/photo"
	"mygram-api/models"
)

type photoMapper struct {
}

func NewPhotoMapper() *photoMapper {
	return &photoMapper{}
}

func (m *photoMapper) ToPhotoResponse(request *models.Photo) *photo.PhotoResponse {
	return &photo.PhotoResponse{
		ID:       request.ID,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
	}
}

func (m *photoMapper) ToPhotoResponses(requests *[]models.Photo) []photo.PhotoResponse {
	var responses []photo.PhotoResponse

	for _, request := range *requests {
		response := m.ToPhotoResponse(&request)
		responses = append(responses, *response)
	}

	if len(responses) > 0 {
		return responses
	}
	return []photo.PhotoResponse{}
}
