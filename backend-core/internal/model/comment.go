package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `gorm:"type:text;not null" json:"content"`

	// Relaciones según tu diagrama SQL
	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	PostID uint `gorm:"not null" json:"post_id"`
	Post   Post `gorm:"foreignKey:PostID" json:"-"` // Ocultamos el objeto Post completo para no hacer ruido

	// Auto-relación (La flecha curva en tu diagrama)
	ParentID *uint     `json:"parent_id,omitempty"`
	Parent   *Comment  `gorm:"foreignKey:ParentID" json:"-"`
	Replies  []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`

	// Reacciones (Opcional: contador simple o relación polimórfica)
	LikesCount int `gorm:"default:0" json:"likes_count"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
