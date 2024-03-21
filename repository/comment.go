package repository

import (
	"errors"
	"mygram-api/domain/requests/comment"
	"mygram-api/models"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(userId uint, request comment.CreateCommentRequest) (*models.Comment, error) {
	var commentModel models.Comment

	db := r.db.Model(&commentModel)

	commentModel.Message = request.Message
	commentModel.PhotoID = request.PhotoID
	commentModel.UserID = userId

	result := db.Debug().Create(&commentModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &commentModel, nil
}

func (r *commentRepository) GetCommentAll() (*[]models.Comment, error) {
	var comments []models.Comment

	db := r.db.Model(&comments)

	result := db.Debug().Preload("User").Preload("Photo").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}

func (r *commentRepository) GetCommentById(commentId uint) (*models.Comment, error) {
	var comment models.Comment

	db := r.db.Model(&comment)

	result := db.Debug().Preload("User").Preload("Photo").Where("id = ?", commentId).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func (r *commentRepository) UpdateComment(userId uint, commentId uint, updatedComment comment.UpdateCommentRequest) (*models.Comment, error) {
	var comment models.Comment

	db := r.db.Model(&comment)

	checkCommentById := db.Debug().Where("id = ?", commentId).First(&comment)
	if checkCommentById.RowsAffected > 1 {
		return &comment, errors.New("comment not found")
	}

	comment.Message = updatedComment.Message

	updateComment := db.Debug().Updates(&comment)
	if updateComment.Error != nil {
		return &comment, updateComment.Error
	}

	return &comment, nil
}

func (r *commentRepository) DeleteComment(commentId uint) (*models.Comment, error) {
	var comment models.Comment

	db := r.db.Model(&comment)

	checkCommentById := db.Debug().Where("id = ?", commentId).First(&comment)
	if checkCommentById.RowsAffected < 1 {
		return &comment, errors.New("comment not found")
	}

	deleteComment := db.Debug().Delete(&comment)
	if deleteComment.Error != nil {
		return &comment, deleteComment.Error
	}

	return &comment, nil
}
