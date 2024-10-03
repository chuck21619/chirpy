package main

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling get chirp")

	chirpIDValue := r.PathValue("chirpID")
	fmt.Println("chirpID: ", chirpIDValue)

	chirpUUID, err := uuid.Parse(chirpIDValue)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't encode parameters", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirp", err)
		return
	}
	
	respondWithJSON(w, http.StatusOK, chirp)
}
