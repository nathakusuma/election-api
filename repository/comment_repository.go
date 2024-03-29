package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (repo *CommentRepository) FindByPostId(postId uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := repo.db.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *CommentRepository) Add(comment *entity.Comment) error {
	return repo.db.Create(&comment).Error
}

func (repo *CommentRepository) FindById(id uint) (*entity.Comment, error) {
	var comment entity.Comment
	if err := repo.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (repo *CommentRepository) Delete(comment *entity.Comment) error {
	return repo.db.Delete(comment).Error
}
