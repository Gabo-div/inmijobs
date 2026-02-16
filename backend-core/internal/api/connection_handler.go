package api

import (
	"encoding/json"
	"net/http"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
	"github.com/go-chi/chi/v5"
)

type ConnectionHandler struct {
	repo *repository.ConnectionRepository
}

func NewConnectionHandler(repo *repository.ConnectionRepository) *ConnectionHandler {
	return &ConnectionHandler{repo: repo}
}

func (h *ConnectionHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connection handler is alive!"))
}



func (h *ConnectionHandler) CreateConnection(w http.ResponseWriter, r *http.Request) {
	var conn model.Connection
	if err := json.NewDecoder(r.Body).Decode(&conn); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&conn); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ConnectionHandler) UpdateConnection(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body struct {
		Status model.ConnectionStatus `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	if err := h.repo.UpdateStatus(id, body.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ConnectionHandler) DeleteConnection(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}