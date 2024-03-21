package service

import (
	socialmedia "mygram-api/domain/requests/social_media"
	socmedRes "mygram-api/domain/responses/social_media"
	"mygram-api/mapper"
	"mygram-api/repository"
)

type socialMediaService struct {
	repository repository.SocialMediaRepository
	mapper     mapper.SocialMediaMapping
}

func NewSocialMediaService(repository repository.SocialMediaRepository, mapper mapper.SocialMediaMapping) *socialMediaService {
	return &socialMediaService{repository, mapper}
}

func (s *socialMediaService) CreateSocialMedia(userId uint, request socialmedia.CreateSocialMediaRequest) (*socmedRes.SocialMediaResponse, error) {
	res, err := s.repository.CreateSocialMedia(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToSocialMediaResponse(res)

	return mapper, nil
}

func (s *socialMediaService) GetSocialMediaAll(userId uint) (*[]socmedRes.SocialMediaWithRelationResponse, error) {
	res, err := s.repository.GetSocialMediaAll(userId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToSocialMediaWithRelationResponses(res)

	return &mapper, nil
}

func (s *socialMediaService) GetSocialMediaById(socialMediaId uint) (*socmedRes.SocialMediaWithRelationResponse, error) {
	res, err := s.repository.GetSocialMediaById(socialMediaId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToSocialMediaWithRelationResponse(res)

	return mapper, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaId uint, request socialmedia.UpdateSocialMediaRequest) (*socmedRes.SocialMediaResponse, error) {
	res, err := s.repository.UpdateSocialMedia(socialMediaId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToSocialMediaResponse(res)

	return mapper, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId uint) (*socmedRes.SocialMediaResponse, error) {
	res, err := s.repository.DeleteSocialMedia(socialMediaId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToSocialMediaResponse(res)

	return mapper, nil
}
