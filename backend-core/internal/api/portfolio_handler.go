package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"github.com/go-chi/chi/v5"
)

type PortfolioHandler struct {
	portfolioService core.PortfolioService
	authService      core.AuthService
}

func NewPortfolioHandler(ps core.PortfolioService, as core.AuthService) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioService: ps,
		authService:      as,
	}
}

func (h *PortfolioHandler) Upload(w http.ResponseWriter, r *http.Request) {
	user, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err == core.ErrUnauthorized {
		utils.RespondError(w, http.StatusUnauthorized, "Authentication required")
		return
	}
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Limitar tamaño del body a 22 MB (overhead del multipart)
	r.Body = http.MaxBytesReader(w, r.Body, 22*1024*1024)
	if err := r.ParseMultipartForm(22 * 1024 * 1024); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Request too large or invalid multipart form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "No file provided. Use field name 'file'")
		return
	}
	defer file.Close()

	portfolioFile, err := h.portfolioService.UploadFile(r.Context(), user.ID, file, header)
	if err == core.ErrInvalidFileType {
		utils.RespondError(w, http.StatusBadRequest, "Invalid file type. Allowed: pdf, doc, docx, jpg, jpeg, png, gif, webp, xlsx, xls, pptx, ppt")
		return
	}
	if err == core.ErrFileTooLarge {
		utils.RespondError(w, http.StatusBadRequest, "File exceeds the maximum allowed size of 20 MB")
		return
	}
	if err != nil {
		slog.Error("[PortfolioHandler][Upload] upload failed", "error", err)
		utils.RespondError(w, http.StatusInternalServerError, "Failed to upload file")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, dto.UploadPortfolioFileResponse{
		Message: "File uploaded successfully",
		File:    portfolioFileToDTO(portfolioFile, r),
	})
}

func (h *PortfolioHandler) GetMyPortfolio(w http.ResponseWriter, r *http.Request) {
	user, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err == core.ErrUnauthorized {
		utils.RespondError(w, http.StatusUnauthorized, "Authentication required")
		return
	}
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	files, err := h.portfolioService.GetUserPortfolio(r.Context(), user.ID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch portfolio")
		return
	}

	fileDTOs := make([]dto.PortfolioFileDTO, len(files))
	for i, f := range files {
		fileDTOs[i] = portfolioFileToDTO(&f, r)
	}

	utils.RespondJSON(w, http.StatusOK, fileDTOs)
}

func (h *PortfolioHandler) GetUserPortfolio(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	files, err := h.portfolioService.GetUserPortfolio(r.Context(), userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch portfolio")
		return
	}

	fileDTOs := make([]dto.PortfolioFileDTO, len(files))
	for i, f := range files {
		fileDTOs[i] = portfolioFileToDTO(&f, r)
	}

	utils.RespondJSON(w, http.StatusOK, fileDTOs)
}

func (h *PortfolioHandler) Download(w http.ResponseWriter, r *http.Request) {
	fileID := chi.URLParam(r, "fileID")

	file, err := h.portfolioService.GetFileForDownload(r.Context(), fileID)
	if err == core.ErrFileNotFound {
		utils.RespondError(w, http.StatusNotFound, "File not found")
		return
	}
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve file")
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.OriginalName))
	w.Header().Set("Content-Type", file.FileType)
	http.ServeFile(w, r, file.StoragePath)
}

func (h *PortfolioHandler) Delete(w http.ResponseWriter, r *http.Request) {
	user, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err == core.ErrUnauthorized {
		utils.RespondError(w, http.StatusUnauthorized, "Authentication required")
		return
	}
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	fileID := chi.URLParam(r, "fileID")

	err = h.portfolioService.DeleteFile(r.Context(), fileID, user.ID)
	if err == core.ErrFileNotFound {
		utils.RespondError(w, http.StatusNotFound, "File not found")
		return
	}
	if err == core.ErrUnauthorizedAccess {
		utils.RespondError(w, http.StatusForbidden, "You don't have permission to delete this file")
		return
	}
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete file")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "File deleted successfully"})
}

// portfolioFileToDTO convierte un model.PortfolioFile en PortfolioFileDTO,
// construyendo la URL pública de descarga dinámicamente a partir de la request.
func portfolioFileToDTO(f *model.PortfolioFile, r *http.Request) dto.PortfolioFileDTO {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	downloadURL := fmt.Sprintf("%s://%s/api/portfolio/files/%s/download", scheme, r.Host, f.ID)

	return dto.PortfolioFileDTO{
		ID:           f.ID,
		OriginalName: f.OriginalName,
		FileType:     f.FileType,
		FileSize:     f.FileSize,
		DownloadURL:  downloadURL,
		CreatedAt:    f.CreatedAt.Time().Unix(),
	}
}
