package dto

import "time"

// CreateCommentReq: Lo que recibes del Frontend para CREAR
type CreateCommentReq struct {
	Content  string `json:"content" binding:"required"` // 'required' lanza error 400 si falta
	ParentID *uint  `json:"parent_id"`                  // Puntero: puede ser null (comentario raíz) o un número (respuesta)
}

// UpdateCommentReq: Lo que recibes del Frontend para EDITAR
type UpdateCommentReq struct {
	Content string `json:"content" binding:"required"`
}

// -------------------------------------------------------------------
// OPCIONAL: Estructuras de Respuesta (Para tener un JSON limpio)
// -------------------------------------------------------------------

// CommentResponse: Lo que envías al Frontend
type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	PostID    uint      `json:"post_id"`
	ParentID  *uint     `json:"parent_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	// Datos del autor (para no enviar el objeto User completo con password, etc.)
	User UserShortResponse `json:"user"`

	// Recursividad: Una lista de comentarios dentro de un comentario
	Replies []CommentResponse `json:"replies,omitempty"`
}

// UserShortResponse: Un resumen del usuario para mostrar en los comentarios
type UserShortResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"` // O "username" según tu modelo User
	Avatar   string `json:"avatar,omitempty"`
}
