package core

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

type CommentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(cr repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: cr}
}

func (s CommentService) ListComments(ctx context.Context, limit int) ([]model.Comment, error) {
	return s.commentRepository.List(ctx, limit)
}

func (s CommentService) GetCommentByID(ctx context.Context, id string) (model.Comment, error) {
	return s.commentRepository.GetByID(ctx, id)
}

func (s CommentService) CreateComment(ctx context.Context, createCommentRequest dto.CreateCommentRequest) (dto.CommentResponse, error) {

	commentModel := model.Comment{
		Content: createCommentRequest.Content,
		UserID:  createCommentRequest.UserID,
		PostID:  createCommentRequest.PostID,
	}
	return s.commentRepository.Create(ctx, commentModel)
}

func (s CommentService) UpdateComment(ctx context.Context, id string, updateCommentRequest dto.UpdateCommentRequest) (model.Comment, error) {

	return s.commentRepository.Update(ctx, id, updateCommentRequest)
}

func (s CommentService) DeleteComment(ctx context.Context, id string) error {
	return s.commentRepository.Delete(ctx, id)
}
