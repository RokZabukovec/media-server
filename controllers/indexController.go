package controllers

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/hub-ui/dist/index.html")
}

func Index(path string) http.Handler {
	fileServer := http.FileServer(http.Dir(path))
	return http.StripPrefix("/", fileServer)
}
