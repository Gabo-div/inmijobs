package core

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(ar repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: ar,
	}
}

// Estructura para leer la respuesta de better-auth
type BetterAuthSessionResponse struct {
	User struct {
		Id string `json:"id"`
	} `json:"user"`
}

func (s AuthService) UserFromHeader(ctx context.Context, header http.Header) (model.User, error) {
	// 1. Extraer el token (del Header o de la Cookie)
	var sessionToken string
	
	authHeader := header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		sessionToken = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		req := &http.Request{Header: header}
		cookie, err := req.Cookie("better-auth.session_token")
		if err == nil && cookie != nil {
			sessionToken = cookie.Value
		}
	}

	if sessionToken == "" {
		slog.Error("[AuthService][UserFromHeader] no session token found")
		return model.User{}, ErrUnauthorized
	}

	// 2. Preguntarle a better-auth si el token es válido
	authReq, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3000/api/auth/get-session", nil)
	if err != nil {
		slog.Error("[AuthService][UserFromHeader] error creating request to auth server", "error", err)
		return model.User{}, ErrUnauthorized
	}
	
	// Le pasamos la cookie al servidor de better-auth
	authReq.Header.Set("cookie", "better-auth.session_token="+sessionToken)

	client := &http.Client{}
	resp, err := client.Do(authReq)
	if err != nil || resp.StatusCode != http.StatusOK {
		slog.Error("[AuthService][UserFromHeader] invalid session according to auth server", "status", resp.StatusCode)
		return model.User{}, ErrUnauthorized
	}
	defer resp.Body.Close()

	var authData BetterAuthSessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&authData); err != nil {
		slog.Error("[AuthService][UserFromHeader] error decoding auth response", "error", err)
		return model.User{}, ErrUnauthorized
	}

	userID := authData.User.Id
	if userID == "" {
		slog.Error("[AuthService][UserFromHeader] missing user id in auth response")
		return model.User{}, ErrUnauthorized
	}

	// 3. Buscar el usuario en la BD de Go con el ID que nos dio better-auth
	user, err := s.authRepository.GetUserById(ctx, userID)
	if err != nil {
		slog.Error("[AuthService][UserFromHeader] unable to get user by id", "error", err, "user_id", userID)
		return model.User{}, ErrUnauthorized
	}

	return user, nil
}