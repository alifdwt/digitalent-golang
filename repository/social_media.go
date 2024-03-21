package repository

import (
	"errors"
	socialmedia "mygram-api/domain/requests/social_media"
	"mygram-api/models"

	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db: db}
}

func (r *socialMediaRepository) CreateSocialMedia(userId uint, request socialmedia.CreateSocialMediaRequest) (*models.SocialMedia, error) {
	var socialMediaModel models.SocialMedia

	db := r.db.Model(&socialMediaModel)

	socialMediaModel.Name = request.Name
	socialMediaModel.SocialMediaURL = request.SocialMediaURL
	socialMediaModel.UserID = userId

	result := db.Debug().Create(&socialMediaModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &socialMediaModel, nil
}

func (r *socialMediaRepository) GetSocialMediaAll(userId uint) (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia

	db := r.db.Model(&socialMedias)

	result := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&socialMedias)
	if result.Error != nil {
		return nil, result.Error
	}

	return &socialMedias, nil
}

func (r *socialMediaRepository) GetSocialMediaById(socialMediaId uint) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	db := r.db.Model(&socialMedia)

	result := db.Debug().Preload("User").Where("id = ?", socialMediaId).First(&socialMedia)
	if result.Error != nil {
		return nil, result.Error
	}

	return &socialMedia, nil
}

func (r *socialMediaRepository) UpdateSocialMedia(socialMediaId uint, request socialmedia.UpdateSocialMediaRequest) (*models.SocialMedia, error) {
	var socialMediaModel models.SocialMedia

	db := r.db.Model(&socialMediaModel)

	checkSocialMediaById := db.Debug().Where("id = ?", socialMediaId).First(&socialMediaModel)
	if checkSocialMediaById.RowsAffected < 1 {
		return &socialMediaModel, errors.New("social media not found")
	}

	socialMediaModel.Name = request.Name
	socialMediaModel.SocialMediaURL = request.SocialMediaURL

	updateSocialMedia := db.Debug().Updates(&socialMediaModel)
	if updateSocialMedia.Error != nil {
		return &socialMediaModel, updateSocialMedia.Error
	}

	return &socialMediaModel, nil
}

func (r *socialMediaRepository) DeleteSocialMedia(socialMediaId uint) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia

	db := r.db.Model(&socialMedia)

	checkSocialMediaById := db.Debug().Where("id = ?", socialMediaId).First(&socialMedia)
	if checkSocialMediaById.RowsAffected < 1 {
		return &socialMedia, errors.New("social media not found")
	}

	deleteSocialMedia := db.Debug().Delete(&socialMedia)
	if deleteSocialMedia.Error != nil {
		return &socialMedia, deleteSocialMedia.Error
	}

	return &socialMedia, nil
}
