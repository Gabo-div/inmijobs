package dto

type PortfolioFileDTO struct {
	ID           string `json:"id"`
	OriginalName string `json:"original_name"`
	FileType     string `json:"file_type"`
	FileSize     int64  `json:"file_size"`
	DownloadURL  string `json:"download_url"`
	CreatedAt    int64  `json:"created_at"`
}

type UploadPortfolioFileResponse struct {
	Message string           `json:"message"`
	File    PortfolioFileDTO `json:"file"`
}
