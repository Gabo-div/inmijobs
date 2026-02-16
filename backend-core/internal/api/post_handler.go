package api

import (
	"net/http"

	"encoding/json"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	svc core.PostService
}

func NewPostHandler(svc core.PostService) *PostHandler {
	return &PostHandler{
		svc: svc,
	}
}

func (h *PostHandler) EditPost(c *gin.Context) {

	var input model.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c.Writer, http.StatusBadRequest, "Datos de entrada inválidos: "+err.Error())
		return
	}

	updatedPost, err := h.svc.UpdatePost(c, input)
	if err != nil {
		utils.RespondError(c.Writer, http.StatusInternalServerError, "Error al actualizar el post: "+err.Error())
		return
	}

	utils.RespondJSON(c.Writer, http.StatusOK, updatedPost)
}

//
// type PostHandler struct {
// 	postService core.PostService
// 	authService core.AuthService
// }
//
// func NewPostHandler(ps core.PostService, as core.AuthService) *PostHandler {
// 	return &PostHandler{
// 		postService: ps,
// 		authService: as,
// 	}
// }

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
