package dto

import "github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"

type CreateJobRequest struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Location       string  `json:"location"`
	CompanyID      string  `json:"company_id"`
	SalaryMin      *int    `json:"salary_min"`
	SalaryMax      *int    `json:"salary_max"`
	EmploymentType string  `json:"employment_type"`
	IsActive       *bool   `json:"is_active"`
	Status         *string `json:"status"`
}

type UpdateJobRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Location       string `json:"location"`
	SalaryMin      *int   `json:"salary_min"`
	SalaryMax      *int   `json:"salary_max"`
	EmploymentType string `json:"employment_type"`
	IsActive       bool   `json:"is_active"`
}

func (r *UpdateJobRequest) ToModel() *model.Job {
	return &model.Job{
		Title:          r.Title,
		Description:    r.Description,
		Location:       r.Location,
		SalaryMin:      r.SalaryMin,
		SalaryMax:      r.SalaryMax,
		EmploymentType: r.EmploymentType,
		IsActive:       r.IsActive,
	}
}
