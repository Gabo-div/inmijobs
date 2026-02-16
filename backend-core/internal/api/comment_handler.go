package api

import (
	"encoding/json"
	"net/http"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type CommentHandler struct {
	commentService core.CommentService
	authService    core.AuthService
}

func NewCommentHandler(cs core.CommentService, as core.AuthService) *CommentHandler {
	return &CommentHandler{
		commentService: cs,
		authService:    as,
	}
}

func (h *CommentHandler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.List)
	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)

	return r
}

func (h *CommentHandler) List(w http.ResponseWriter, r *http.Request) {
	comments, err := h.commentService.ListComments(r.Context(), 100)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "unable to list comments: "+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, err := h.commentService.GetCommentByID(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "comment not found: "+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, comment)
}

func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {

	user, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized: "+err.Error())
		return
	}

	var comment dto.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	comment.UserID = user.ID

	created, err := h.commentService.CreateComment(r.Context(), comment)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "unable to create comment: "+err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusCreated, created)
}

func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "id")

	_, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized: "+err.Error())
		return
	}

	var commentRequest dto.UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentRequest); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	updated, err := h.commentService.UpdateComment(r.Context(), commentID, commentRequest)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "unable to update comment: "+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, updated)
}

func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized: "+err.Error())
		return
	}

	if err := h.commentService.DeleteComment(r.Context(), id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "unable to delete comment: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
