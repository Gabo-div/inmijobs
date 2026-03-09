package core

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"log/slog"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"gorm.io/gorm"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
)

type PostService interface {
	UpdatePost(ctx context.Context, postID string, req dto.CreatePostRequest) (model.Post, error)
	CreatePost(ctx context.Context, req dto.CreatePostRequest) (*model.Post, error)
	GetByID(ctx context.Context, id string) (*dto.PostResponseDTO, error)
	DeletePost(ctx context.Context, id string) (*model.Post, error)
	GetFeed(ctx context.Context, userID string, limit int, cursor string) ([]model.Post, *string, error)
	GetImagesByUserID(ctx context.Context, userID string) ([]string, error)
}

type postService struct {
	repo    repository.PostRepo
	jobRepo repository.JobRepository
}

func NewPostService(repo repository.PostRepo, jobRepo repository.JobRepository) PostService {
	return &postService{
		repo:    repo,
		jobRepo: jobRepo,
	}
}

func (s *postService) CreatePost(ctx context.Context, req dto.CreatePostRequest) (*model.Post, error) {

	if req.Title == "" || req.Content == "" {
		return nil, errors.New("Write a title!")
	}

	if req.JobID != nil && req.JobID == nil {
		return nil, errors.New("invalid job id")
	}

	var postImages []model.Image
	for _, url := range req.Images {
		if url != "" {
			postImages = append(postImages, model.Image{
				URL: url,
			})
		}
	}
	post := model.Post{
		ID:        utils.NewID(),
		Title:     req.Title,
		UserID:    req.UserID,
		JobID:     req.JobID,
		CompanyID: req.CompanyID,
		Content:   req.Content,
		Images:    postImages,
	}
	if post.JobID != nil {
		job, err := s.jobRepo.GetJobByID(ctx, *post.JobID)
		if err != nil {
			slog.Error("[PostService] Job no encontrado", "jobID", post.JobID, "error", err)
			return nil, fmt.Errorf("la vacante con ID %d no existe", post.JobID)
		}

		post.JobID = &job.ID
		post.CompanyID = &job.CompanyID

		if job.IsActive != true {
			return nil, fmt.Errorf("no se puede publicar: la vacante está %s", job.Status)
		} else if req.CompanyID != nil {

			post.CompanyID = req.CompanyID
		}

	}
	err := s.repo.CreatePost(ctx, &post)

	if err != nil {
		slog.Error("[PostService][CreatePost] unable create post", "error", err)
		return nil, err
	}
	return &post, nil
}

func (s *postService) UpdatePost(ctx context.Context, postID string, req dto.CreatePostRequest) (model.Post, error) {

	existingPost, err := s.repo.GetByID(ctx, postID)
	if err != nil {
		return model.Post{}, errors.New("post no encontrado")
	}

	p := model.Post{
		ID:        postID,
		Title:     req.Title,
		Content:   req.Content,
		UserID:    existingPost.UserID,
		JobID:     req.JobID,
		CompanyID: req.CompanyID,
		CreatedAt: existingPost.CreatedAt,
	}

	if len(req.Images) > 0 {
		for _, url := range req.Images {
			p.Images = append(p.Images, model.Image{
				URL: url,
			})
		}
	}

	return s.repo.EditPost(ctx, postID, p)
}
func (s *postService) GetByID(ctx context.Context, id string) (*dto.PostResponseDTO, error) {

	post, err := s.repo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("The post not found")
		}
		return nil, err
	}
	response := utils.MapToCleanPost(post)

	return response, nil
}

func (s *postService) DeletePost(ctx context.Context, id string) (*model.Post, error) {

	if s.repo.IsAlreadyDeleted(ctx, id) {
		slog.Error("The post is already Deleted")
		return nil, errors.New("The post is already deleted")
	}
	post, err := s.repo.DeletePost(ctx, id)

	if err != nil {
		slog.Error("[PostService][CreatePost] Unable to delete post", "error", err)
		return nil, err
	}
	return post, nil
}

func splitCursor(cursor string) (string, string) {
	if cursor == "" {
		return "", ""
	}
	parts := strings.SplitN(cursor, "_", 2)

	if len(parts) != 2 {
		return "", ""
	}
	// parts[0] is the timestamp, parts[1] is the ID
	return parts[0], parts[1]
}

func (s *postService) GetFeed(ctx context.Context, userID string, limit int, cursor string) ([]model.Post, *string, error) {
	var rawCursor string
	if cursor != "" {
		decodedBytes, err := base64.StdEncoding.DecodeString(cursor)
		if err != nil {
			return nil, nil, errors.New("Invalid cursor")
		}
		rawCursor = string(decodedBytes)
	}
	createdAt, postID := splitCursor(rawCursor)

	posts, err := s.repo.GetFeed(ctx, userID, limit, createdAt, postID)
	if err != nil {
		return nil, nil, err
	}
	var nextCursor *string

	if len(posts) > 0 {
		lastPost := posts[len(posts)-1]
		formattedTime := lastPost.CreatedAt.Format("2006-01-02 15:04:05")
		rawCursor := fmt.Sprintf("%s_%s", formattedTime, lastPost.ID)

		encoded := base64.StdEncoding.EncodeToString([]byte(rawCursor))
		nextCursor = &encoded
	}

	return posts, nextCursor, nil
}

func (s *postService) GetImagesByUserID(ctx context.Context, userID string) ([]string, error) {
	return s.repo.GetImagesByUserID(ctx, userID)
}
