package api

import (
    "net/http"

    "github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
    "github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
    "github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
)

type PingHandler struct {
    authService core.AuthService
}

func NewPingHandler(as core.AuthService) *PingHandler {
    return &PingHandler{
        authService: as,
    }
}

func (h PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
    user, err := h.authService.UserFromHeader(r.Context(), r.Header)

    if err == core.ErrUnauthorized {
        utils.RespondJSON(w, http.StatusOK, dto.Ping{
            Message:       "Pong",
            Authenticated: false,
        })
        return
    }

    if err != nil {
        utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
        return
    }

    utils.RespondJSON(w, http.StatusOK, dto.Ping{
        Message:       "Pong",
        Authenticated: true,
        UserID:        user.ID,
    })
}