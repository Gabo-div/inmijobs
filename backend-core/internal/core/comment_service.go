package core

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

type CommentServ interface {
	ListComments(ctx context.Context, limit int) ([]model.Comment, error)
	GetCommentByID(ctx context.Context, id string) (model.Comment, error)
	CreateComment(ctx context.Context, createCommentRequest dto.CreateCommentRequest) (dto.CommentResponse, error)
	UpdateComment(ctx context.Context, id string, updateCommentRequest dto.UpdateCommentRequest) (model.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}

type commentService struct {
	commentRepository repository.CommentRepo
}

func NewCommentService(cr repository.CommentRepo) CommentServ {
	return &commentService{commentRepository: cr}
}

func (s commentService) ListComments(ctx context.Context, limit int) ([]model.Comment, error) {
	return s.commentRepository.List(ctx, limit)
}

func (s commentService) GetCommentByID(ctx context.Context, id string) (model.Comment, error) {
	return s.commentRepository.GetByID(ctx, id)
}

func (s commentService) CreateComment(ctx context.Context, createCommentRequest dto.CreateCommentRequest) (dto.CommentResponse, error) {

	commentModel := model.Comment{
		Content: createCommentRequest.Content,
		UserID:  createCommentRequest.UserID,
		PostID:  createCommentRequest.PostID,
	}
	return s.commentRepository.Create(ctx, commentModel)
}

func (s commentService) UpdateComment(ctx context.Context, id string, updateCommentRequest dto.UpdateCommentRequest) (model.Comment, error) {

	return s.commentRepository.Update(ctx, id, updateCommentRequest)
}

func (s commentService) DeleteComment(ctx context.Context, id string) error {
	return s.commentRepository.Delete(ctx, id)
}
