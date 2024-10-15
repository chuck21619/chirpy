package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chuck21619/chirpy/internal/auth"
	"github.com/chuck21619/chirpy/internal/database"
)

func (cfg *apiConfig) revokeRefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("revoking refresh token")
	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		return
	}

	refreshToken, err := cfg.db.GetRefreshTokenFromRefreshToken(r.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		return
	}

	revokeParams := database.RevokeParams{
		UserID: refreshToken.UserID,
		UpdatedAt: time.Now(),
	}
	
	err = cfg.db.Revoke(r.Context(), revokeParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}