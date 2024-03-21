package service

import (
	"mygram-api/domain/requests/photo"
	"mygram-api/mapper"
	"mygram-api/repository"

	photoRes "mygram-api/domain/responses/photo"
)

type photoService struct {
	repository repository.PhotoRepository
	mapper     mapper.PhotoMapping
}

func NewPhotoService(repository repository.PhotoRepository, mapper mapper.PhotoMapping) *photoService {
	return &photoService{
		repository: repository,
		mapper:     mapper,
	}
}

func (s *photoService) CreatePhoto(userId uint, request photo.CreatePhotoRequest) (*photoRes.PhotoResponse, error) {
	res, err := s.repository.CreatePhoto(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhoto(res)

	return mapper, nil
}

func (s *photoService) GetPhotoAll() (*[]photoRes.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponses(res)

	return &mapper, nil
}

func (s *photoService) GetPhotoById(photoId uint) (*photoRes.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoById(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}

func (s *photoService) UpdatePhoto(userId uint, photoId uint, request photo.UpdatePhotoRequest) (*photoRes.PhotoResponse, error) {
	res, err := s.repository.UpdatePhoto(userId, photoId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhoto(res)

	return mapper, nil
}

func (s *photoService) DeletePhoto(photoId uint) (*photoRes.PhotoResponse, error) {
	res, err := s.repository.DeletePhoto(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhoto(res)

	return mapper, nil
}
