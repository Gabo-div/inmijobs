package model

type Rol string

const (
	RolAdmin      Rol = "Admin"
	RolEmployment Rol = "Employment"
	RolUser       Rol = "User"
)

type Post struct {
	ID           uint          `gorm:"primaryKey"`
	JobID        *int          `gorm:"not null"`
	Job          *Job          `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID       int           `gorm:"not null"`
	User         User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyID    int           `gorm:"not null"`
	Company      *Company      `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content      string        `gorm:"not null"`
	Created_at   UnixTime      `gorm:"not null;autoCreateTime"`
	Update_at    UnixTime      `gorm:"not null;autoCreateTime"`
	Deleted_at   UnixTime      `gorm:"not null;autoCreateTime"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Interactions []Interaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Images       []Image       `gorm:"many2many:PostImage;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Job struct {
	ID          uint    `gorm:"primaryKey"`
	CompanyID   int     `gorm:"not null"`
	Company     Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RecruiterID int     `gorm:"not null"`
	Recruiter   User    `gorm:"foreignKey:RecruiterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Status      string  `gorm:"not null"`
	Posts       []Post  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Image struct {
	ID         uint     `gorm:"primaryKey"`
	Caption    string   `gorm:"not null"`
	URL        string   `gorm:"not null"`
	Created_at UnixTime `gorm:"not null;autoCreateTime"`
}

type Comment struct {
	ID         uint     `gorm:"primaryKey"`
	UserID     int      `gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostID     int      `gorm:"not null"`
	Post       Post     `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content    string   `gorm:"type:text"`
	Created_at UnixTime `gorm:"not null;autoCreateTime"`
	Update_at  UnixTime `gorm:"autoUpdateTime"`
	Deleted_at UnixTime `gorm:"not null;autoDeleteTime"`
}

type Company struct {
	ID       uint   `gorm:"primaryKey"`
	OwnerID  uint   `gorm:"index"`
	Owner    User   `gorm:"foreignKey:OwnerID"`
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
}

type Reaction struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	IconURL string `gorm:"not null"`
}

type Interaction struct {
	ID         uint     `gorm:"primaryKey"`
	UserID     int      `gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID"`
	PostID     int      `gorm:"not null"`
	Post       Post     `gorm:"foreignKey:PostID"`
	ReactionID int      `gorm:"not null;uniqueIndex"`
	Reaction   Reaction `gorm:"foreignKey:ReactionID"`
	Created_at UnixTime `gorm:"not null;autoCreateTime"`
}

type Employee struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    int     `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	CompanyID int     `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
	Rol       Rol     `gorm:"index;type:Roel;not null"`
}

type PostImage struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	URL  string `gorm:"not null"`
}
