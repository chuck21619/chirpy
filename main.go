package main

import (
	"net/http"
)

func main() {

	multiplexer := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: multiplexer,
	}
	
	
	filesDirectory := http.Dir(".")
	fileServerHandler := http.FileServer(filesDirectory)
	strippedPrefixFileServerHandler := http.StripPrefix("/app", fileServerHandler)
	multiplexer.Handle("/app/", strippedPrefixFileServerHandler)
	//multiplexer.Handle("/", fileServerHandler)

	multiplexer.HandleFunc("/healthz", mysWEetnest)

	server.ListenAndServe()
}

func mysWEetnest(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "text/plain; charset=utf-8")
	//responseWriter.WriteHeader(http.StatusOK) //unnecessary
	statusString := http.StatusText(http.StatusOK)
	//statusString := "any string i fucking want bitch"
	responseWriter.Write([]byte(statusString))
	
}