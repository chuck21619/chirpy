package main

import (
	"net/http"
	"fmt"
)

func main() {

	multiplexer := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: multiplexer,
	}
	
	apiConfig := apiConfig{
		fileserverHits: 0,
	}
	
	filesDirectory := http.Dir(".")
	fileServerHandler := http.FileServer(filesDirectory)
	strippedPrefixFileServerHandler := http.StripPrefix("/app", fileServerHandler)
	wrappedStrippedFileServerHandler := apiConfig.middlewareMetricsInc(strippedPrefixFileServerHandler)
	multiplexer.Handle("/app/", wrappedStrippedFileServerHandler)
	//multiplexer.Handle("/", fileServerHandler)

	multiplexer.HandleFunc("/healthz", mysWEetnest)

	multiplexer.HandleFunc("/metrics", apiConfig.handlerMetrics)
	multiplexer.HandleFunc("/reset", apiConfig.handlerReset)

	server.ListenAndServe()
}

func mysWEetnest(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "text/plain; charset=utf-8")
	//responseWriter.WriteHeader(http.StatusOK) //unnecessary
	statusString := http.StatusText(http.StatusOK)
	//statusString := "any string i fucking want bitch"
	responseWriter.Write([]byte(statusString))
	
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits)))
}

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}

