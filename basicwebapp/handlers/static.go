package handlers

import (
	"net/http"
)

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	http.ServeFile(w, r, filename)
}
