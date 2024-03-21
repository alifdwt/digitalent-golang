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

func (m *photoMapper) ToPhoto(request *models.Photo) *photo.PhotoResponse {
	return &photo.PhotoResponse{
		ID:       request.ID,
		Caption:  request.Caption,
		Title:    request.Title,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
	}
}

func (m *photoMapper) ToPhotoResponse(request *models.Photo) *photo.PhotoWithRelationResponse {
	return &photo.PhotoWithRelationResponse{
		ID:       request.ID,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
		User:     *NewUserMapper().ToUserRelationsResponse(&request.User),
	}
}

func (m *photoMapper) ToPhotoResponses(requests *[]models.Photo) []photo.PhotoWithRelationResponse {
	var responses []photo.PhotoWithRelationResponse

	for _, request := range *requests {
		response := m.ToPhotoResponse(&request)
		responses = append(responses, *response)
	}

	if len(responses) > 0 {
		return responses
	}
	return []photo.PhotoWithRelationResponse{}
}
