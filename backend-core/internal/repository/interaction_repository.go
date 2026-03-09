package repository

import (
	"errors"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type InteractionRepo interface {
	TogglePostReaction(userID string, postID string, reactionID int) (*model.Interaction, string, error)
	GetReactionsByPost(postID string) ([]model.Interaction, error)
}
type interactionRepository struct {
	db *gorm.DB
}

func NewInteractionRepository(db *gorm.DB) InteractionRepo {
	return &interactionRepository{db: db}
}

func (r *interactionRepository) TogglePostReaction(userID string, postID string, reactionID int) (*model.Interaction, string, error) {

	var existing model.Interaction

	if userID == "" || postID == "" {
		return nil, "", errors.New("el ID de usuario y el ID de post son obligatorios")
	}

	err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&existing).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, "", err
	}

	if err == gorm.ErrRecordNotFound {

		newInteraction := model.Interaction{
			UserID:     userID,
			PostID:     postID,
			ReactionID: reactionID,
		}
		errCreate := r.db.Create(&newInteraction).Error
		return &newInteraction, "created", errCreate
	}

	if existing.ReactionID == reactionID {

		errDelete := r.db.Delete(&existing).Error
		return nil, "deleted", errDelete
	}

	// Existe pero es OTRA reacción: La actualizamos
	existing.ReactionID = reactionID

	if errUpdate := r.db.Save(&existing).Error; errUpdate != nil {
		return nil, "", errUpdate
	}
	return &existing, "updated", nil
}

func (r *interactionRepository) GetReactionsByPost(postID string) ([]model.Interaction, error) {
	var interactions []model.Interaction
	err := r.db.Where("post_id = ?", postID).Find(&interactions).Error
	return interactions, err
}
