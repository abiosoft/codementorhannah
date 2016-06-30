package handlers

import (
	"net/http"
)

func handleErr(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
