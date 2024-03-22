package repository

import (
	"errors"
	"mygram-api/domain/requests/auth"
	"mygram-api/domain/requests/user"
	"mygram-api/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(registerReq *auth.RegisterRequest) (*models.User, error) {
	var user models.User

	user.Username = registerReq.Username
	user.Email = registerReq.Email
	user.Password = registerReq.Password
	user.Age = registerReq.Age
	user.ProfileImageURL = registerReq.ProfileImageURL

	db := r.db.Model(&user)

	result := db.Debug().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email not found")
		}
		return nil, errors.New("error while fetching email: " + err.Error())
	}

	return &user, nil
}

func (r *userRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUserById := db.Debug().Where("id = ?", id).First(&user)
	if checkUserById.RowsAffected < 0 {
		return &user, errors.New("user not found")
	}

	return &user, nil
}

func (r *userRepository) UpdateUserById(id uint, updatedUser *user.UpdateUserRequest) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUserById := db.Debug().Where("id = ?", id).First(&user)
	if checkUserById.RowsAffected > 1 {
		return &user, errors.New("user not found")
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Age = updatedUser.Age
	user.ProfileImageURL = updatedUser.ProfileImageURL

	updateUser := db.Debug().Updates(&user)
	if updateUser.Error != nil {
		return nil, updateUser.Error
	}

	return &user, nil
}

func (r *userRepository) DeleteUserById(id uint) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	res, err := r.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if err := db.Delete(&res).Error; err != nil {
		return nil, errors.New("error while deleting user: " + err.Error())
	}

	return res, nil
}
