package model

type Company struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Weblink     string
	LinkedinURL string
	Number      string
	Description string `gorm:"type:text"`
	Sector      string
	Foundation  string
	Size        string
	Logo        *string
	Banner      *string
	CreatedAt   UnixTime `gorm:"autoCreateTime"`
	UpdatedAt   UnixTime `gorm:"autoUpdateTime"`

	// Relación con dueño
	UserID string `gorm:"not null;index"`
	Owner  User   `gorm:"foreignKey:UserID"`

	// Relaciones
	Jobs []Job `gorm:"foreignKey:CompanyID"`
}

type Job struct {
	ID             string `gorm:"primaryKey"`
	Title          string `gorm:"not null"`
	Description    string `gorm:"type:text;not null"`
	Location       string `gorm:"not null"`
	SalaryMin      *int
	SalaryMax      *int
	EmploymentType string   `gorm:"not null"`
	IsActive       bool     `gorm:"default:true"`
	CreatedAt      UnixTime `gorm:"autoCreateTime"`
	UpdatedAt      UnixTime `gorm:"autoUpdateTime"`

	CompanyID string  `gorm:"not null;index"`
	Company   Company `gorm:"foreignKey:CompanyID"`

	Applications []Application `gorm:"foreignKey:JobID"`
}

type Application struct {
	ID          string   `gorm:"primaryKey"`
	CoverLetter *string  `gorm:"type:text"`
	Status      string   `gorm:"default:'pending'"`
	CreatedAt   UnixTime `gorm:"autoCreateTime"`
	UpdatedAt   UnixTime `gorm:"autoUpdateTime"`

	UserID string `gorm:"not null;index"`
	User   User   `gorm:"foreignKey:UserID"`

	JobID string `gorm:"not null;index"`
	Job   Job    `gorm:"foreignKey:JobID"`
}
