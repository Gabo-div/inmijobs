package model

import (
	"time"

	"gorm.io/gorm"
)

type Rol string

const (
	RolAdmin      Rol = "Admin"
	RolEmployment Rol = "Employment"
	RolUser       Rol = "User"
)

type Post struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Title        string `gorm:"type:text;not null"`
	Content      string `gorm:"type:text;not null"`
	UserID       string   `gorm:"not null"`
	User         User   `gorm:"foreignKey:UserID"`
	JobID        *int
	Job          Job `gorm:"foreignKey:JobID"`
	CompanyID    *int
	Company      Company        `gorm:"foreignKey:CompanyID"`
	Comments     []Comment      `gorm:"foreignKey:PostID"`
	Interactions []Interaction  `gorm:"foreignKey:PostID"`
	Images       []Image        `gorm:"many2many:post_images;"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"` 
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Job struct {
	ID          int     `gorm:"primaryKey"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Status      string  `gorm:"not null"`
	CompanyID   int     `gorm:"not null"`
	Company     Company `gorm:"foreignKey:CompanyID"`
	RecruiterID string     `gorm:"not null" json:"recruiter_id"`
	Recruiter   User    `gorm:"foreignKey:RecruiterID"`

	Posts []Post  `gorm:"foreignKey:JobID"`
}

type Company struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
	OwnerID  string   `gorm:"index"`
	Owner    User   `gorm:"foreignKey:OwnerID"`
}

type Employee struct {
	ID        int     `gorm:"primaryKey"`
	UserID    string     `gorm:"not null" `
	User      User    `gorm:"foreignKey:UserID"`
	CompanyID int     `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
	Rol       Rol     `gorm:"index;type:rol;not null"`
}

type Interaction struct {
	ID         int       `gorm:"primaryKey"`
	UserID     string    `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	PostID     *uint     // Modificado: Puntero (*) para permitir nulos
	Post       Post      `gorm:"foreignKey:PostID"`
	CommentID  *int      // NUEVO: Puntero para tu tarea de comentarios
	ReactionID int       `gorm:"not null"`
	Reaction   Reaction  `gorm:"foreignKey:ReactionID"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Reaction struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	IconURL string `gorm:"not null"`
}

type Image struct {
	ID        int      `gorm:"primaryKey"`
	Name      string   `gorm:"type:text"`
	URL       string   `gorm:"not null"`
	Caption   string   `gorm:"type:text" json:"caption"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type PostImage struct {
	PostID    int      `gorm:"primaryKey"`
	ImageID   uint     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
