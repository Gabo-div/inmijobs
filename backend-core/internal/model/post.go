package model

import (
	"time"

	"gorm.io/gorm"
)

// --- ENUMS Y CONSTANTES ---

type Rol string

const (
	RolAdmin      Rol = "Admin"
	RolEmployment Rol = "Employment"
	RolUser       Rol = "User"
)

// --- STRUCTS PRINCIPALES ---

type Post struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	// Relaciones
	UserID uint  `json:"user_id"`
	User   *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	JobID *int `json:"job_id,omitempty"`
	Job   *Job `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"job,omitempty"`

	CompanyID *int     `json:"company_id,omitempty"`
	Company   *Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company,omitempty"`

	// Listas (Has Many / Many to Many)
	Comments     []Comment     `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
	Interactions []Interaction `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"interactions,omitempty"`
	Images       []Image       `gorm:"many2many:post_images;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images,omitempty"`

	// Tiempos
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Job struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Status      string `gorm:"not null" json:"status"`

	CompanyID int      `gorm:"not null" json:"company_id"`
	Company   *Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company,omitempty"`

	RecruiterID int   `gorm:"not null" json:"recruiter_id"`
	Recruiter   *User `gorm:"foreignKey:RecruiterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"recruiter,omitempty"`

	Posts []Post `gorm:"foreignKey:JobID" json:"posts,omitempty"`
}

type Company struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Location string `gorm:"not null" json:"location"`

	OwnerID uint  `gorm:"index" json:"owner_id"`
	Owner   *User `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
}

type Employee struct {
	ID     uint  `gorm:"primaryKey" json:"id"`
	UserID int   `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	CompanyID int      `gorm:"not null" json:"company_id"`
	Company   *Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`

	Rol Rol `gorm:"index;type:varchar(50);not null" json:"rol"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `gorm:"type:text" json:"content"`

	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	PostID uint  `gorm:"not null" json:"post_id"`
	Post   *Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Interaction struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	PostID uint  `gorm:"not null" json:"post_id"`
	Post   *Post `gorm:"foreignKey:PostID" json:"-"`

	ReactionID int       `gorm:"not null" json:"reaction_id"`
	Reaction   *Reaction `gorm:"foreignKey:ReactionID" json:"reaction,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

type Reaction struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	IconURL string `gorm:"not null" json:"icon_url"`
}

type Image struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	URL       string    `gorm:"not null" json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

// Esta tabla intermedia es necesaria para la relación many2many de Post <-> Image
type PostImage struct {
	PostID  uint `gorm:"primaryKey"`
	ImageID uint `gorm:"primaryKey"`
}
