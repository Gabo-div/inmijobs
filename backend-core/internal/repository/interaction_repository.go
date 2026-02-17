package repository

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type InteractionRepository struct {
	db *gorm.DB
}

func NewInteractionRepository(db *gorm.DB) *InteractionRepository {
	return &InteractionRepository{db: db}
}

// TogglePostReaction maneja agregar, quitar o cambiar la reacción de un POST
func (r *InteractionRepository) TogglePostReaction(userID string, postID uint, reactionID int) (*model.Interaction, string, error) {
	

	var existing model.Interaction

	// Buscamos si ya existe una interacción de este usuario en este post
	err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		// No existe: La creamos
		newInteraction := model.Interaction{
			UserID:     userID,
			PostID:     &postID,
			ReactionID: reactionID,
		}
		errCreate := r.db.Create(&newInteraction).Error
		return &newInteraction, "created", errCreate
	}

	if existing.ReactionID == reactionID {
		// Existe y es la MISMA reacción: La eliminamos (Quitar Like)
		errDelete := r.db.Delete(&existing).Error
		return nil, "deleted", errDelete
	}

	// Existe pero es OTRA reacción: La actualizamos
	existing.ReactionID = reactionID
	errUpdate := r.db.Save(&existing).Error
	return &existing, "updated", errUpdate
}

// GetReactionsByPost obtiene todas las reacciones de un post específico
func (r *InteractionRepository) GetReactionsByPost(postID uint) ([]model.Interaction, error) {
	var interactions []model.Interaction
	err := r.db.Where("post_id = ?", postID).Find(&interactions).Error
	return interactions, err
}