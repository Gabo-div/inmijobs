package repository

import (
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"gorm.io/gorm"
)

type ConnectionRepository struct {
	db *gorm.DB
}

func NewConnectionRepository(db *gorm.DB) *ConnectionRepository {
	return &ConnectionRepository{db: db}
}

func (r *ConnectionRepository) Create(conn *model.Connection) error {
	return r.db.Create(conn).Error
}

func (r *ConnectionRepository) UpdateStatus(id string, status model.ConnectionStatus) error {
	return r.db.Model(&model.Connection{}).Where("id = ?", id).Update("status", status).Error
}

func (r *ConnectionRepository) Delete(id string) error {
	return r.db.Delete(&model.Connection{}, id).Error
}