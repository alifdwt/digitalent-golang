package service

import (
	socialmedia "mygram-api/domain/requests/social_media"
	"mygram-api/models"
	"mygram-api/repository"
)

type socialMediaService struct {
	repository repository.SocialMediaRepository
}

func NewSocialMediaService(repository repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{repository}
}

func (s *socialMediaService) CreateSocialMedia(userId uint, request socialmedia.CreateSocialMediaRequest) (*models.SocialMedia, error) {
	res, err := s.repository.CreateSocialMedia(userId, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *socialMediaService) GetSocialMediaAll(userId uint) (*[]models.SocialMedia, error) {
	res, err := s.repository.GetSocialMediaAll(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *socialMediaService) GetSocialMediaById(socialMediaId uint) (*models.SocialMedia, error) {
	res, err := s.repository.GetSocialMediaById(socialMediaId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaId uint, request socialmedia.UpdateSocialMediaRequest) (*models.SocialMedia, error) {
	res, err := s.repository.UpdateSocialMedia(socialMediaId, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId uint) (*models.SocialMedia, error) {
	res, err := s.repository.DeleteSocialMedia(socialMediaId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
