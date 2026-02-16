package dto

type CreateCommentRequest struct {
	Content string `json:"message"`
	UserID  string `json:"userId,omitempty"`
	PostID  string `json:"postId,omitempty"`
}

type UpdateCommentRequest struct {
	Content string `json:"message"`
	UserID  string `json:"userId,omitempty"`
}

type CommentResponse struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	PostID    string `json:"postId"`
	UserID    string `json:"userId"`
	CreatedAt int64  `json:"createdAt"` // timestamp Unix
	UpdatedAt int64  `json:"updatedAt"` // timestamp Unix
}
