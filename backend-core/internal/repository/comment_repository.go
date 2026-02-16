package repository

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model" // ⚠️ Ajusta esto a tu módulo real (ej: github.com/AnDoor/inmijobs/...)
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

// Create inserta un nuevo comentario
func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.DB.Create(comment).Error
}

// GetByPostID busca comentarios de un post.
// Preload("User"): Carga la info del autor.
// Preload("Replies"): Carga las respuestas directas (si definiste la relación en el modelo).
func (r *CommentRepository) GetByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment

	// Truco: Traemos solo los comentarios "Padre" (ParentID IS NULL) y sus respuestas.
	// Si prefieres traer TODO plano y armar el árbol en el front, quita el .Where("parent_id IS NULL")
	err := r.DB.Where("post_id = ? AND parent_id IS NULL", postID).
		Preload("User").         // Cargar autor del comentario padre
		Preload("Replies").      // Cargar respuestas
		Preload("Replies.User"). // Cargar autor de las respuestas
		Order("created_at desc").
		Find(&comments).Error

	return comments, err
}

// Update actualiza el contenido
func (r *CommentRepository) Update(id uint, content string) (*model.Comment, error) {
	var comment model.Comment
	// Primero buscamos si existe
	if err := r.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}

	comment.Content = content
	err := r.DB.Save(&comment).Error
	return &comment, err
}

// Delete elimina un comentario
func (r *CommentRepository) Delete(id uint) error {
	// Unscoped() elimina permanentemente. Quítalo si usas SoftDelete.
	return r.DB.Unscoped().Delete(&model.Comment{}, id).Error
}
