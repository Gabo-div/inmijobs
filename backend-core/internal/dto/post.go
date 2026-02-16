package dto

type CreatePostRequest struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	JobID     *int   `json:"job_id,omitempty"`
	CompanyID *int   `json:"company_id,omitempty"`
}
