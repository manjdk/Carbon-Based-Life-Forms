package controller

import (
	"net/http"
)

func NewReDocHandler(rootDir string) http.HandlerFunc {
	handler := http.StripPrefix("/api/", http.FileServer(http.Dir(rootDir)))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		handler.ServeHTTP(w, r)
	}
}

func NewSwaggerHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}
