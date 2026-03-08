package dto

type CreateCommentRequest struct {
	Content string `json:"message"`
	UserID  string `json:"user_id,omitempty"`
	PostID  string `json:"post_id,omitempty"`
}

type UpdateCommentRequest struct {
	Content string `json:"message"`
	UserID  string `json:"user_id,omitempty"`
}

type CommentResponse struct {
	ID        uint `json:"id"`
	Message   string `json:"message"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	CreatedAt int64  `json:"createdAt"` // timestamp Unix
	UpdatedAt int64  `json:"updatedAt"` // timestamp Unix
}
