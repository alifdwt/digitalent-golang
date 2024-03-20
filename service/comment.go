package service

import (
	"mygram-api/domain/requests/comment"
	"mygram-api/mapper"
	"mygram-api/repository"

	commentRes "mygram-api/domain/responses/comment"
)

type commentService struct {
	repository repository.CommentRepository
	mapper     mapper.CommentMapping
}

func NewCommentService(repository repository.CommentRepository, mapper mapper.CommentMapping) *commentService {
	return &commentService{repository, mapper}
}

func (s *commentService) CreateComment(userId uint, request comment.CreateCommentRequest) (*commentRes.CommentResponse, error) {
	res, err := s.repository.CreateComment(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}

func (s *commentService) GetCommentAll() (*[]commentRes.CommentResponse, error) {
	res, err := s.repository.GetCommentAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponses(res)

	return &mapper, nil
}

func (s *commentService) GetCommentById(commentId uint) (*commentRes.CommentResponse, error) {
	res, err := s.repository.GetCommentById(commentId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}

func (s *commentService) UpdateComment(userId uint, commentId uint, request comment.UpdateCommentRequest) (*commentRes.CommentResponse, error) {
	res, err := s.repository.UpdateComment(userId, commentId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}

func (s *commentService) DeleteComment(commentId uint) (*commentRes.CommentResponse, error) {
	res, err := s.repository.DeleteComment(commentId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}
