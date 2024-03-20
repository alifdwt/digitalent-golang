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

func (m *commentMapper) ToCommentResponses(requests *[]models.Comment) []comment.CommentResponse {
	var responses []comment.CommentResponse
	for _, request := range *requests {
		response := m.ToCommentResponse(&request)
		responses = append(responses, *response)
	}

	if len(responses) > 0 {
		return responses
	}
	return []comment.CommentResponse{}
}
