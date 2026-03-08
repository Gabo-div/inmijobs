package model

import "time"

type Interaction struct {
	ID         int       `gorm:"primaryKey"`
	UserID     string    `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	PostID     string    
	Post       Post      `gorm:"foreignKey:PostID"`
	CommentID  *int      
	ReactionID int       `gorm:"not null"`
	Reaction   Reaction  `gorm:"foreignKey:ReactionID"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Reaction struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	IconURL string `gorm:"not null"`
}
