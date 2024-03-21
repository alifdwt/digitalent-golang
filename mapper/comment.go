package mapper

import (
	"mygram-api/domain/responses/comment"
	"mygram-api/models"
)

type commentMapper struct {
}

func NewCommentMapper() *commentMapper {
	return &commentMapper{}
}

func (m *commentMapper) ToCommentResponse(request *models.Comment) *comment.CommentResponse {
	return &comment.CommentResponse{
		ID:      request.ID,
		Message: request.Message,
		UserID:  request.UserID,
		PhotoID: request.PhotoID,
	}
}

func (m *commentMapper) ToCommentWithRelationResponse(request *models.Comment) *comment.CommentWithRelationResponse {
	return &comment.CommentWithRelationResponse{
		ID:      request.ID,
		Message: request.Message,
		PhotoID: request.PhotoID,
		UserID:  request.UserID,
		User:    *NewUserMapper().ToUserRelationsResponse(&request.User),
		Photo:   *NewPhotoMapper().ToPhoto(&request.Photo),
	}
}

func (m *commentMapper) ToCommentWithRelationResponses(requests *[]models.Comment) []comment.CommentWithRelationResponse {
	var responses []comment.CommentWithRelationResponse
	for _, request := range *requests {
		response := m.ToCommentWithRelationResponse(&request)
		responses = append(responses, *response)
	}

	if len(responses) > 0 {
		return responses
	}
	return []comment.CommentWithRelationResponse{}
}
