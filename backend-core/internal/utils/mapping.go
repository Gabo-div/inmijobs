package utils

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
)

func MapToCleanPost(post *model.Post) *dto.PostResponseDTO {
	imageURLs := make([]string, len(post.Images))
	for i, img := range post.Images {
		imageURLs[i] = img.URL
	}

	interactionsDTO := make([]dto.InteractionShortDTO, len(post.Interactions))
	for i, inter := range post.Interactions {
		interactionsDTO[i] = dto.InteractionShortDTO{
			ID:       inter.ID,
			UserName: inter.User.Name,
			Reaction: inter.Reaction.Name,
		}
	}


	commentsDTO := make([]dto.CommentShortDTO, len(post.Comments))
	for i, c := range post.Comments {
		
		commentsDTO[i] = dto.CommentShortDTO{
			ID:        c.ID,
			Content:   c.Content,
			UserName:  c.User.Name,
			CreatedAt: c.CreatedAt,
		}
	}


	response := &dto.PostResponseDTO{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		User: dto.UserShortDTO{
			ID:   post.User.ID,
			Name: post.User.Name,
		},
		Images:       imageURLs,
		Interactions: interactionsDTO,
		Comments:     commentsDTO,
	}

	
	if post.JobID != nil && &post.Job.ID != nil {
		response.Job = &dto.JobShortDTO{
			ID:    post.Job.ID,
			Title: post.Job.Title,
		}
	}


	if post.CompanyID != nil && &post.Company.ID != nil {
		response.Company = &dto.CompanyShortDTO{
			ID:       post.Company.ID,
			Name:     post.Company.Name,
			Location: post.Company.Location,
		}
	}

	return response
}
