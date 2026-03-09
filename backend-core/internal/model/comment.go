package model

type Comment struct {
	ID        uint `gorm:"primaryKey; autoIncrement"`
	Content   string `gorm:"not null"`
	PostID    string `gorm:"not null"`
	CreatedAt int64  `gorm:"not null;autoCreateTime"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime"`

	UserID string `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`

	// Relaciones adicionales
	// Post    Post    `gorm:"foreignKey:PostID;references:ID"`
}

