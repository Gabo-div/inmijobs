package repository

import (
	"context"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type PortfolioRepository struct {
	db gorm.DB
}

func NewPortfolioRepository(db gorm.DB) *PortfolioRepository {
	return &PortfolioRepository{db: db}
}

func (r *PortfolioRepository) CreatePortfolioFile(ctx context.Context, file *model.PortfolioFile) error {
	return r.db.WithContext(ctx).Create(file).Error
}

func (r *PortfolioRepository) GetPortfolioFilesByUserID(ctx context.Context, userID string) ([]model.PortfolioFile, error) {
	var files []model.PortfolioFile
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&files).Error
	return files, err
}

func (r *PortfolioRepository) GetPortfolioFileByID(ctx context.Context, fileID string) (*model.PortfolioFile, error) {
	var file model.PortfolioFile
	if err := r.db.WithContext(ctx).First(&file, "id = ?", fileID).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *PortfolioRepository) DeletePortfolioFile(ctx context.Context, fileID string) error {
	return r.db.WithContext(ctx).Delete(&model.PortfolioFile{}, "id = ?", fileID).Error
}
