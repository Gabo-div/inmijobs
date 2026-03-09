package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID             string `gorm:"primaryKey"`
	Title   string `gorm:"type:text;not null"`
	Content string `gorm:"type:text;not null"`

	UserID string `gorm:"index"`
	User   User   `gorm:"foreignKey:UserID"`

	JobID *string `gorm:"index"`
	Job   *Job    `gorm:"foreignKey:JobID"`

	CompanyID *string  `gorm:"index"`
	Company   *Company `gorm:"foreignKey:CompanyID"`

	Comments     []Comment      `gorm:"foreignKey:PostID"`
	Interactions []Interaction  `gorm:"foreignKey:PostID"`
	Images       []Image        `gorm:"many2many:post_images;"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Image struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:text"`
	URL       string    `gorm:"not null"`
	Caption   string    `gorm:"type:text" json:"caption"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type PostImage struct {
	PostID    string    `gorm:"primaryKey"`
	ImageID   uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
