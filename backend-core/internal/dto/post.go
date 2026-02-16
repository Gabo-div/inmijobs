package dto

type CreatePostRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	JobID     *int   `json:"job_id,omitempty"`
	CompanyID *int   `json:"company_id,omitempty"`
}
