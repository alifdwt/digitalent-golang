package repository

import (
	"errors"
	"mygram-api/domain/requests/photo"
	"mygram-api/models"

	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) CreatePhoto(userId uint, request photo.CreatePhotoRequest) (*models.Photo, error) {
	var photoModel models.Photo

	db := r.db.Model(&photoModel)

	photoModel.Title = request.Title
	photoModel.Caption = request.Caption
	photoModel.PhotoURL = request.PhotoURL
	photoModel.UserID = userId

	result := db.Debug().Create(&photoModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photoModel, nil
}

func (r *photoRepository) GetPhotoAll() (*[]models.Photo, error) {
	var photos []models.Photo

	db := r.db.Model(&photos)

	result := db.Debug().Preload("User").Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photos, nil
}

func (r *photoRepository) GetPhotoById(photoId uint) (*models.Photo, error) {
	var photo models.Photo

	db := r.db.Model(&photo)

	result := db.Debug().Preload("User").Where("id = ?", photoId).First(&photo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photo, nil
}

func (r *photoRepository) UpdatePhoto(userId uint, photoId uint, updatedPhoto photo.UpdatePhotoRequest) (*models.Photo, error) {
	var photo models.Photo

	db := r.db.Model(&photo)

	checkPhotoById := db.Debug().Where("id = ? AND user_id = ?", photoId, userId).First(&photo)
	if checkPhotoById.RowsAffected > 1 {
		return &photo, errors.New("photo not found")
	}

	photo.Title = updatedPhoto.Title
	photo.PhotoURL = updatedPhoto.PhotoURL
	photo.Caption = updatedPhoto.Caption

	updatePhoto := db.Debug().Updates(&photo)
	if updatePhoto.Error != nil {
		return &photo, updatePhoto.Error
	}

	return &photo, nil
}

func (r *photoRepository) DeletePhoto(photoId uint) (*models.Photo, error) {
	var photo models.Photo

	db := r.db.Model(&photo)

	checkPhotoById := db.Debug().Where("id = ?", photoId).First(&photo)
	if checkPhotoById.RowsAffected < 1 {
		return &photo, errors.New("photo not found")
	}

	deletePhoto := db.Debug().Delete(&photo)
	if deletePhoto.Error != nil {
		return &photo, deletePhoto.Error
	}

	return &photo, nil
}
