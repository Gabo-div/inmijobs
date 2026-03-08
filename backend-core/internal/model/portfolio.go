package model

type PortfolioFile struct {
	ID           string   `gorm:"primaryKey"`
	UserID       string   `gorm:"not null;index"`
	User         User     `gorm:"foreignKey:UserID"`
	FileName     string   `gorm:"not null"`        // nombre único en disco
	OriginalName string   `gorm:"not null"`        // nombre original del archivo
	FileType     string   `gorm:"not null"`        // MIME type
	FileSize     int64    `gorm:"not null"`        // tamaño en bytes
	StoragePath  string   `gorm:"not null;unique"` // ruta relativa en disco
	CreatedAt    UnixTime `gorm:"autoCreateTime"`
	UpdatedAt    UnixTime `gorm:"autoUpdateTime"`
}
