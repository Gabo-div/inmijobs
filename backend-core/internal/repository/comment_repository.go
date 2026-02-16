package repository

import (
	"context"
	"fmt"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db gorm.DB
}

func NewCommentRepository(db gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r CommentRepository) Create(ctx context.Context, comment model.Comment) (dto.CommentResponse, error) {
	if err := gorm.G[model.Comment](&r.db).Create(ctx, &comment); err != nil {
		return dto.CommentResponse{}, fmt.Errorf("repository: create comment: %w", err)
	}

	resp := dto.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Content,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	return resp, nil
}

func (r CommentRepository) Update(ctx context.Context, CommentID string, comment dto.UpdateCommentRequest) (model.Comment, error) {

	var existing model.Comment
	if _, err := gorm.G[model.Comment](&r.db).Where("id = ?", CommentID).First(ctx); err != nil {
		return model.Comment{}, fmt.Errorf("repository: find comment before update: %w", err)
	}

	if comment.Content != "" {
		existing.Content = comment.Content
	}
	if comment.Content != "" {
		existing.Content = comment.Content
	}

	_, err := gorm.G[model.Comment](&r.db).Where("id = ?", CommentID).Updates(ctx, existing)
	if err != nil {
		return model.Comment{}, fmt.Errorf("repository: update comment: %w", err)
	}

	return existing, nil
}

func (r CommentRepository) Delete(ctx context.Context, id string) error {
	_, err := gorm.G[model.Comment](&r.db).
		Where("id = ?", id).
		Delete(ctx)

	if err != nil {
		return fmt.Errorf("repository: delete comment: %w", err)
	}

	return nil
}

func (r CommentRepository) GetByID(ctx context.Context, id string) (model.Comment, error) {
	comment, err := gorm.G[model.Comment](&r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return model.Comment{}, fmt.Errorf("repository: get comment by id: %w", err)
	}
	return comment, nil
}

func (r CommentRepository) List(ctx context.Context, limit int) ([]model.Comment, error) {

	comments, err := gorm.G[model.Comment](&r.db).
		Order("created_at desc").
		Limit(limit).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("repository: list comments: %w", err)
	}

	return comments, nil
}
