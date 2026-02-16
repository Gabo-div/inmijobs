package core

import (
	"context"
	// "errors"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
	"github.com/gin-gonic/gin"
	"log/slog"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
)

type PostService interface {
	UpdatePost(ctx *gin.Context, input model.Post) (model.Post, error)
	CreatePost(ctx context.Context, req dto.CreatePostRequest) error
}

type postService struct {
	repo repository.PostRepo
}

func NewPostService(repo repository.PostRepo) PostService {
	return &postService{
		repo: repo,
	}
}

func (s *postService) CreatePost(ctx context.Context, req dto.CreatePostRequest) error {

	post := model.Post{
		// ID:        utils.NewID(),
		Title:     req.Title,
		UserID:    req.UserID,
		JobID:     req.JobID,
		CompanyID: req.CompanyID,
		Content:   req.Content,
	}

	err := s.repo.CreatePost(ctx, &post)

	if err != nil {
		slog.Error("[PostService][CreatePost] unable create post", "error", err)
		return err
	}
	return nil
}

func (s *postService) UpdatePost(ctx *gin.Context, input model.Post) (model.Post, error) {

	updatedPost, err := s.repo.EditPost(*ctx, input)
	if err != nil {
		return model.Post{}, err
	}

	return updatedPost, nil
}
