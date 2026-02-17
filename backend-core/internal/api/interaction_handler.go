package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/go-chi/chi/v5"
)

type InteractionHandler struct {
	service *core.InteractionService
}

func NewInteractionHandler(service *core.InteractionService) *InteractionHandler {
	return &InteractionHandler{service: service}
}

// POST /posts/{id}/reactions
func (h *InteractionHandler) TogglePostReaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Obtener ID del post desde la URL usando Chi
	postIDStr := chi.URLParam(r, "id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "ID de post inválido"}`, http.StatusBadRequest)
		return
	}

	// 2. Decodificar el JSON que envía el cliente
	var req dto.ReactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "JSON inválido o incompleto"}`, http.StatusBadRequest)
		return
	}

	// 3. Llamar a tu servicio (notar que casteamos a uint y dejamos reactionID como int)
	interaction, action, err := h.service.TogglePostReaction(req.UserID, uint(postID), req.ReactionID)
	if err != nil {
		http.Error(w, `{"error": "Error procesando la reacción"}`, http.StatusInternalServerError)
		return
	}

	// 4. Preparar la respuesta
	if action == "deleted" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Reacción eliminada (Like quitado)"})
		return
	}

	response := dto.ReactionResponse{
		InteractionID: interaction.ID,
		UserID:        interaction.UserID,
		ReactionID:    interaction.ReactionID,
		Action:        action,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Reacción procesada con éxito",
		"data":    response,
	})
}

// GET /posts/{id}/reactions
func (h *InteractionHandler) GetPostReactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postIDStr := chi.URLParam(r, "id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "ID de post inválido"}`, http.StatusBadRequest)
		return
	}

	interactions, err := h.service.GetReactionsByPost(uint(postID))
	if err != nil {
		http.Error(w, `{"error": "Error obteniendo reacciones"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"data": interactions})
}