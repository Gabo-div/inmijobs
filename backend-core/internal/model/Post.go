package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"` // Asumo campos básicos
	Content string `json:"content"`

	UserID uint `json:"user_id"`
	// Si ya tienes el modelo User definido:
	// User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// Relación inversa con comentarios (opcional, pero útil)
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
