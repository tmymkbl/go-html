package controllers

import (
	"net/http"
)

func e401(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "401")
}
func e404(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "404")
}
func e500(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "500")
}
