package repository

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type PostRepo interface {
	EditPost(ctx context.Context, postID int, p model.Post) (model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepo {
	return &PostRepository{db: db}
}

func (r *PostRepository) EditPost(ctx context.Context, postID int, p model.Post) (model.Post, error) {

	var editedPost model.Post

	if err := r.db.WithContext(ctx).Model(&editedPost).Where("id = ?", postID).Updates(&p).Error; err != nil {
		return model.Post{}, err
	}

	if len(p.Comments) > 0 {
		r.db.Model(&editedPost).Association("Comments").Replace(p.Comments)
	}
	if len(p.Interactions) > 0 {
		r.db.Model(&editedPost).Association("Interactions").Replace(p.Interactions)
	}
	if len(p.Images) > 0 {
		r.db.Model(&editedPost).Association("Images").Replace(p.Images)
	}

	return editedPost, r.db.WithContext(ctx).
		Preload("Comments").
		Preload("Interactions").
		Preload("Images").
		First(&editedPost, postID).Error

}

func (r *PostRepository) CreatePost(ctx context.Context, post *model.Post) error {
	err := gorm.G[model.Post](r.db).Create(ctx, post)
	if err != nil {
		return err
	}
	return nil
}
