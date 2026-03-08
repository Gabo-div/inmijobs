package repository

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepo interface {
	EditPost(ctx context.Context, postID string, p model.Post) (model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) error
	GetByID(ctx context.Context, id string) (*model.Post, error)
	DeletePost(ctx context.Context, id string) (*model.Post, error)
	IsAlreadyDeleted(ctx context.Context, id string) bool
	GetFeed(ctx context.Context, userID string, limit int, createdAt string, postID string) ([]model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepo {
	return &postRepository{db: db}
}

func (r *postRepository) GetByID(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post

	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Company").
		Preload("Company").
		Preload("Job").
		Preload("Images").
		Preload("Interactions").
		Preload("Interactions.User").
		Preload("Interactions.Reaction").
		Preload("Comments.User").
		Preload("Comments").
		First(&post, "id = ?", id).Error

	return &post, err
}

func (r *postRepository) EditPost(ctx context.Context, postID string, p model.Post) (model.Post, error) {
	var editedPost model.Post

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&editedPost, "posts.id = ?", postID).Error; err != nil {
			return err
		}

		if err := tx.Model(&editedPost).Updates(p).Error; err != nil {
			return err
		}

		if len(p.Images) > 0 {

			if err := tx.Model(&editedPost).Association("Images").Replace(p.Images); err != nil {
				return err
			}
		}

		if len(p.Comments) > 0 {
			if err := tx.Model(&editedPost).Association("Comments").Replace(p.Comments); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return model.Post{}, err
	}

	res, err := r.GetByID(ctx, postID)
	if err != nil {
		return model.Post{}, err
	}
	return *res, nil
}
func (r *postRepository) CreatePost(ctx context.Context, post *model.Post) error {
	if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (r *postRepository) DeletePost(ctx context.Context, id string) (*model.Post, error) {
	post := model.Post{ID: id}

	if err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Delete(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) IsAlreadyDeleted(ctx context.Context, id string) bool {
	post := model.Post{ID: id}

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return true
	}
	return false
}

func (r *postRepository) GetFeed(ctx context.Context, userID string, limit int, createdAt string, postID string) ([]model.Post, error) {
	var posts []model.Post

	sub1 := r.db.Table("connections").
		Select("requester_id").
		Where("receiver_id = ? AND status = ?", userID, "accepted")

	sub2 := r.db.Table("connections").
		Select("receiver_id").
		Where("requester_id = ? AND status = ?", userID, "accepted")

	query := r.db.Table("posts").
		Where("(user_id IN (?) OR user_id IN (?) OR user_id = ?)", sub1, sub2, userID)

	if createdAt != "" && postID != "" {
		query = query.Where("(posts.created_at < ?) OR (posts.created_at = ? AND posts.id < ?)",
			createdAt, createdAt, postID)
	}

	err := query.Order("posts.created_at DESC, posts.id DESC").
		Limit(limit).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}
