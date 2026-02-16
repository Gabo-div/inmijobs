package model

type Rol string

const (
	RolAdmin      Rol = "Admin"
	RolEmployment Rol = "Employment"
	RolUser       Rol = "User"
)

type Post struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"type:text;not null"`
	Content      string `gorm:"type:text;not null"`
	UserID       uint
	User         User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	JobID        *int
	Job          Job `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyID    *int
	Company      Company       `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Interactions []Interaction `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Images       []Image       `gorm:"many2many:post_images;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    UnixTime      `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    UnixTime      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt    *UnixTime     `gorm:"index;column:deleted_at"`
}

type Job struct {
	ID          uint    `gorm:"primaryKey"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Status      string  `gorm:"not null"`
	CompanyID   int     `gorm:"not null"`
	Company     Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RecruiterID int     `gorm:"not null" json:"recruiter_id"`
	Recruiter   User    `gorm:"foreignKey:RecruiterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Posts []Post `gorm:"foreignKey:JobID" json:"posts,omitempty"`
}

type Company struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
	OwnerID  uint   `gorm:"index"`
	Owner    User   `gorm:"foreignKey:OwnerID"`
}

type Employee struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    int     `gorm:"not null" `
	User      User    `gorm:"foreignKey:UserID"`
	CompanyID int     `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
	Rol       Rol     `gorm:"index;type:rol;not null"`
}

type Interaction struct {
	ID         uint     `gorm:"primaryKey"`
	UserID     uint     `gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID"`
	PostID     uint     `gorm:"not null"`
	Post       Post     `gorm:"foreignKey:PostID"`
	ReactionID int      `gorm:"not null"`
	Reaction   Reaction `gorm:"foreignKey:ReactionID"`
	CreatedAt  UnixTime `gorm:"not null;autoCreateTime"`
}

type Reaction struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	IconURL string `gorm:"not null"`
}

type Image struct {
	ID        uint     `gorm:"primaryKey"`
	Name      string   `gorm:"not null"`
	URL       string   `gorm:"not null"`
	CreatedAt UnixTime `gorm:"not null;autoCreateTime"`
}

type PostImage struct {
	PostID    uint     `gorm:"primaryKey"`
	ImageID   uint     `gorm:"primaryKey"`
	CreatedAt UnixTime `gorm:"not null;autoCreateTime"`
}
