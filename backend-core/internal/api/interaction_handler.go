package api

import (
	"encoding/json"
	"net/http"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type InteractionHandler struct {
	service core.InteractionRepo
}

func NewInteractionHandler(service core.InteractionRepo) *InteractionHandler {
	return &InteractionHandler{service: service}
}

func (h *InteractionHandler) TogglePostReaction(w http.ResponseWriter, r *http.Request) {

	Postid := chi.URLParam(r, "id")
	if Postid == "" {
		utils.RespondError(w, http.StatusBadRequest, "Missing post ID")
		return
	}

	var req dto.ReactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "ERROR: Invalid or incomplete JSON")
		return
	}

	interaction, action, err := h.service.TogglePostReaction(req.UserID, Postid, req.ReactionID)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "ERROR: Loading reaction's info")
		return
	}

	if action == "deleted" {
		utils.RespondJSON(w, http.StatusAccepted, "The reaction was deleted successfully")

		return
	}

	response := dto.ReactionResponse{
		InteractionID: interaction.ID,
		UserID:        interaction.UserID,
		ReactionID:    interaction.ReactionID,
		Action:        action,
	}
	utils.RespondJSON(w, http.StatusCreated, response)
}

func (h *InteractionHandler) GetPostReactions(w http.ResponseWriter, r *http.Request) {

	Postid := chi.URLParam(r, "id")
	if Postid == "" {
		utils.RespondError(w, http.StatusBadRequest, "Missing post ID")
		return
	}

	interactions, err := h.service.GetReactionsByPost(Postid)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "ERROR: error listing reactions")
		return
	}
	utils.RespondJSON(w, http.StatusOK, interactions)
}
