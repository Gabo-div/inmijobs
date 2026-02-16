package api

import (
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	svc core.PostService
}

func NewPostHandler(svc core.PostService) *PostHandler {
	return &PostHandler{
		svc: svc,
	}
}

func (h *PostHandler) EditPost(w http.ResponseWriter, r *http.Request) {

	var input model.Post
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondError(w, http.StatusBadRequest, "Missing post ID")
		return
	}

	Postid, err := strconv.Atoi(id)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid post ID format")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedPost, err := h.svc.UpdatePost(r.Context(), Postid, input)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create post: "+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedPost)
}

func (p PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	// user, err := p.authService.UserFromHeader(r.Context(), r.Header)
	// if err != nil {
	// 	utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }

	var req dto.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	req.UserID = 123 //Should be ---> &user.ID

	err := p.svc.CreatePost(r.Context(), req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	utils.RespondJSON(w, http.StatusOK, dto.CreatePostRequest{
		UserID:    req.UserID,
		CompanyID: req.CompanyID,
		Content:   req.Content,
	})
}
