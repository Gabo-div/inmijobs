package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
)

var (
	ErrInvalidFileType = errors.New("invalid file type")
	ErrFileTooLarge    = errors.New("file too large")
	ErrFileNotFound    = errors.New("file not found")
)

const (
	maxFileSize    = 20 * 1024 * 1024 // 20 MB
	uploadBasePath = "uploads/portfolio"
)

var allowedExtensions = map[string]bool{
	".pdf":  true,
	".doc":  true,
	".docx": true,
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
	".xlsx": true,
	".xls":  true,
	".pptx": true,
	".ppt":  true,
}

type PortfolioService struct {
	portfolioRepo repository.PortfolioRepository
}

func NewPortfolioService(pr repository.PortfolioRepository) *PortfolioService {
	return &PortfolioService{portfolioRepo: pr}
}

func (s *PortfolioService) UploadFile(ctx context.Context, userID string, file multipart.File, header *multipart.FileHeader) (*model.PortfolioFile, error) {
	// Validar tamaño
	if header.Size > maxFileSize {
		return nil, ErrFileTooLarge
	}

	// Validar extensión
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExtensions[ext] {
		return nil, ErrInvalidFileType
	}

	// Crear directorio si no existe
	if err := os.MkdirAll(uploadBasePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generar nombre único en disco
	fileID := utils.NewID()
	storedName := fileID + ext
	storagePath := filepath.Join(uploadBasePath, storedName)

	// Guardar archivo en disco
	dst, err := os.Create(storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(storagePath)
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	// Detectar MIME type a partir de la extensión
	mimeType := extensionToMIME(ext)

	// Guardar registro en DB
	portfolioFile := &model.PortfolioFile{
		ID:           fileID,
		UserID:       userID,
		FileName:     storedName,
		OriginalName: header.Filename,
		FileType:     mimeType,
		FileSize:     header.Size,
		StoragePath:  storagePath,
	}

	if err := s.portfolioRepo.CreatePortfolioFile(ctx, portfolioFile); err != nil {
		os.Remove(storagePath)
		return nil, fmt.Errorf("failed to save file metadata: %w", err)
	}

	return portfolioFile, nil
}

func (s *PortfolioService) GetUserPortfolio(ctx context.Context, userID string) ([]model.PortfolioFile, error) {
	return s.portfolioRepo.GetPortfolioFilesByUserID(ctx, userID)
}

func (s *PortfolioService) GetFileForDownload(ctx context.Context, fileID string) (*model.PortfolioFile, error) {
	file, err := s.portfolioRepo.GetPortfolioFileByID(ctx, fileID)
	if err != nil {
		return nil, ErrFileNotFound
	}
	return file, nil
}

func (s *PortfolioService) DeleteFile(ctx context.Context, fileID, userID string) error {
	file, err := s.portfolioRepo.GetPortfolioFileByID(ctx, fileID)
	if err != nil {
		return ErrFileNotFound
	}

	if file.UserID != userID {
		return ErrUnauthorizedAccess
	}

	// Eliminar del disco
	os.Remove(file.StoragePath)

	// Eliminar de DB
	return s.portfolioRepo.DeletePortfolioFile(ctx, fileID)
}

// extensionToMIME retorna el MIME type según la extensión
func extensionToMIME(ext string) string {
	mimeMap := map[string]string{
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".webp": "image/webp",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		".xls":  "application/vnd.ms-excel",
		".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		".ppt":  "application/vnd.ms-powerpoint",
	}
	if m, ok := mimeMap[ext]; ok {
		return m
	}
	return "application/octet-stream"
}
