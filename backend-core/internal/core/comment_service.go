package core

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

type CommentService struct {
	repo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

// GetComments: Obtiene la lista de comentarios de un post
func (s *CommentService) GetComments(postID uint) ([]model.Comment, error) {
	return s.repo.GetByPostID(postID)
}

// CreateComment: Crea un comentario o respuesta
func (s *CommentService) CreateComment(userID uint, postID uint, req dto.CreateCommentReq) (*model.Comment, error) {
	newComment := &model.Comment{
		Content:  req.Content,
		UserID:   userID,
		PostID:   postID,
		ParentID: req.ParentID, // Si es nil, es un comentario raíz
	}

	err := s.repo.Create(newComment)
	return newComment, err
}

// UpdateComment: Edita un comentario existente
func (s *CommentService) UpdateComment(id uint, req dto.UpdateCommentReq) (*model.Comment, error) {
	// Aquí podrías agregar validación de usuario (si es dueño del comentario)
	return s.repo.Update(id, req.Content)
}

// DeleteComment: Elimina un comentario
func (s *CommentService) DeleteComment(id uint) error {
	return s.repo.Delete(id)
}
