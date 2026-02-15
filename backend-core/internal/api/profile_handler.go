package api

import (
	"encoding/json"
	"net/http"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/dto"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
)

type ProfileHandler struct {
	profileService core.ProfileService
	authService    core.AuthService
}

func NewProfileHandler(ps core.ProfileService, as core.AuthService) *ProfileHandler {
	return &ProfileHandler{
		profileService: ps,
		authService:    as,
	}
}

func (h ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {

	user, err := h.authService.UserFromHeader(r.Context(), r.Header)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	//user := struct{ ID string }{ID: "user_test_123"}
	var req dto.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	profile, err := h.profileService.UpdateProfile(r.Context(), user.ID, req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	utils.RespondJSON(w, http.StatusOK, dto.ProfileResponse{
		ID:        profile.ID,
		UserID:    profile.UserID,
		Biography: profile.Biography,
		Title:     profile.Title,
		Location:  profile.Location,
	})
}
