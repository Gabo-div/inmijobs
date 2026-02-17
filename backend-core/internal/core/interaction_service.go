package core

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

type InteractionService struct {
	repo *repository.InteractionRepository
}

func NewInteractionService(repo *repository.InteractionRepository) *InteractionService {
	return &InteractionService{repo: repo}
}

func (s *InteractionService) TogglePostReaction(userID string, postID uint, reactionID int) (*model.Interaction, string, error) {
	return s.repo.TogglePostReaction(userID, postID, reactionID)
}

func (s *InteractionService) GetReactionsByPost(postID uint) ([]model.Interaction, error) {
	return s.repo.GetReactionsByPost(postID)
}