package database

import (
	"log/slog"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) {
	// Create Test User
	userID := "user_test_123"
	user := model.User{
		ID:            userID,
		Name:          "Test User",
		Email:         "test@example.com",
		EmailVerified: true,
	}

	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&user)

	if result.Error != nil {
		slog.Error("[Database] Failed to seed user", "error", result.Error)
	} else if result.RowsAffected > 0 {
		slog.Info("[Database] Seeded test user")
	}

	// Create Test Profile
	profileID := utils.NewID()
	profile := model.Profile{
		ID:        profileID,
		UserID:    userID,
		Biography: toPtr("This is a test biography."),
		Location:  toPtr("Test City, Country"),
	}

	result = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoNothing: true,
	}).Create(&profile)

	if result.Error != nil {
		slog.Error("[Database] Failed to seed profile", "error", result.Error)
	} else if result.RowsAffected > 0 {
		slog.Info("[Database] Seeded test profile")
	}

}

func toPtr(s string) *string {
	return &s
}
