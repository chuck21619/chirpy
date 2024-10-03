package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chuck21619/chirpy/internal/database"
	"github.com/google/uuid"
)

func (a *apiConfig) createChirp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("creating chirp")

	type parameters struct {
		Body string `json:"body"`
		User_id string `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(params.Body, badWords)

	uuidParam, err := uuid.Parse(params.User_id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't encode parameters", err)
	}

	createChirpParams := database.CreateChirpParams{
		Body: cleaned,
		UserID: uuidParam,
	}

	chirp, err := a.db.CreateChirp(r.Context(), createChirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, chirp)
}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}