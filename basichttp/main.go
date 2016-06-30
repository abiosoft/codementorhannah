package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/error", handleError)

	println("Listening 8080")
	http.ListenAndServe(":8080", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	successJSON(w, "Hello")
}

func handleError(w http.ResponseWriter, r *http.Request) {
	errorJSON(w, "Unknown error")
}

func successJSON(w io.Writer, body interface{}) {
	// JSend format
	var resp = map[string]interface{}{
		"status": "success",
		"data":   body,
	}
	json.NewEncoder(w).Encode(resp)
}

func errorJSON(w http.ResponseWriter, err string) {
	// JSend format
	var resp = map[string]interface{}{
		"status": "error",
		"error":  err,
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(resp)
}
