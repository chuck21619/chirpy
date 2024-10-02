package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "platform is not dev", nil)
		return
	}

	err := cfg.db.Reset(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error deleting users", nil)
		return
	}
	
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
}