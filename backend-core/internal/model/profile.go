package model

type Profile struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"not null;uniqueIndex"`
	Biography *string
	Title     *string
	Location  *string
	CreatedAt UnixTime `gorm:"not null;autoCreateTime"`
	UpdatedAt UnixTime `gorm:"not null;autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
