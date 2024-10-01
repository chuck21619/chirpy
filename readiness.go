package main

import "net/http"

func ready(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "text/plain; charset=utf-8")
	//responseWriter.WriteHeader(http.StatusOK) //unnecessary 
	statusString := http.StatusText(http.StatusOK)
	//statusString := "any string i fucking want bitch"
	responseWriter.Write([]byte(statusString))
	
}