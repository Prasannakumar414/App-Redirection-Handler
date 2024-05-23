package main

import (
	"log"
	"net/http"
)

/**
Server for handling redirections for testing apps.
*/
func main() {
	log.Println("Intializing Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect", playStoreRedirectionHandler)
	mux.HandleFunc("/.well-known/assetlinks.json", appLinkVerification)
	log.Println("Server is running")
	http.ListenAndServe(":8080", mux)
}
