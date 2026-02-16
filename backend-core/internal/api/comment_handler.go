package api

import (
	"net/http"
	"strconv"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service *core.CommentService
}

func NewCommentHandler(service *core.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// GET /posts/:id/comments
func (h *CommentHandler) GetCommentsByPost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de post inválido"})
		return
	}

	comments, err := h.service.GetComments(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener comentarios"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// POST /posts/:id/comments
func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de post inválido"})
		return
	}

	var req dto.CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// ---------------------------------------------------------------
	// TODO: Aquí debes obtener el ID real del usuario desde el token JWT
	// Por ahora usamos '1' como mock para que puedas probar
	// userID := c.GetUint("userID")
	userID := uint(1)
	// ---------------------------------------------------------------

	comment, err := h.service.CreateComment(userID, uint(postID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el comentario"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// PUT /comments/:id
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de comentario inválido"})
		return
	}

	var req dto.UpdateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateComment(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DELETE /comments/:id
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteComment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comentario eliminado"})
}
