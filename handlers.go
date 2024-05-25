package main

import (
	"log"
	"net/http"
	"strings"
)

/**
For handling redirection to playstore based on the query params.
*/
func playStoreRedirectionHandler(w http.ResponseWriter, r *http.Request) {
	var userAgent = r.UserAgent()
	var values = r.URL.Query()
	var redirectUrl = ""
	log.Println("Handling requirest from " + userAgent)
	if strings.Contains(userAgent, "Android") {
		redirectUrl = "https://play.google.com/store/apps/details?" + values.Encode()
	} else {
		redirectUrl = "https://apxor.com"
	}
	log.Println(redirectUrl)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

/**
Fort handling verification of app links in application.
*/
func appLinkVerification(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling requirest for verification")
	http.ServeFile(w, r, "assetlinks.json")
}
