package mapper

import (
	socialmedia "mygram-api/domain/responses/social_media"
	"mygram-api/models"
)

type socialMediaMapper struct {
}

func NewSocialMediaMapper() *socialMediaMapper {
	return &socialMediaMapper{}
}

func (m *socialMediaMapper) ToSocialMediaResponse(request *models.SocialMedia) *socialmedia.SocialMediaResponse {
	return &socialmedia.SocialMediaResponse{
		ID:             request.ID,
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         request.UserID,
	}
}

func (m *socialMediaMapper) ToSocialMediaWithRelationResponse(request *models.SocialMedia) *socialmedia.SocialMediaWithRelationResponse {
	return &socialmedia.SocialMediaWithRelationResponse{
		ID:             request.ID,
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         request.UserID,
		User:           *NewUserMapper().ToUserRelationsResponse(&request.User),
	}
}

func (m *socialMediaMapper) ToSocialMediaWithRelationResponses(requests *[]models.SocialMedia) []socialmedia.SocialMediaWithRelationResponse {
	var responses []socialmedia.SocialMediaWithRelationResponse
	for _, request := range *requests {
		response := m.ToSocialMediaWithRelationResponse(&request)
		responses = append(responses, *response)
	}

	if len(responses) > 0 {
		return responses
	}

	return []socialmedia.SocialMediaWithRelationResponse{}
}
