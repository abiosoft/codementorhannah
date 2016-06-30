package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	// Auth and Log
	users := rootChain()
	users.Use(func(w http.ResponseWriter, r *http.Request) error {
		body := map[string]interface{}{
			"name":  "Hann",
			"email": "a@a.com",
			"age":   100,
		}
		return json.NewEncoder(w).Encode(body)
	})

	http.Handle("/users", users)

	// Auth and Log
	docs := rootChain()
	docs.Use(func(w http.ResponseWriter, r *http.Request) error {
		body := map[string]interface{}{
			"name":  "Hann",
			"email": "a@a.com",
			"age":   100,
		}
		return json.NewEncoder(w).Encode(body)
	})

	http.Handle("/docs", docs)

	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)

}

var rootChain = func() *CustomChain {
	return &CustomChain{
		[]Middleware{
			authMid,
			logMid,
		},
	}
}

// Middleware is a custom handler
type Middleware func(w http.ResponseWriter, r *http.Request) error

// CustomChain is a middleware chain
type CustomChain struct {
	middlewares []Middleware
}

// Use uses a middleware.
func (c *CustomChain) Use(m Middleware) {
	c.middlewares = append(c.middlewares, m)
}

// ServeHTTP satisfies http.Handler
func (c *CustomChain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, middleware := range c.middlewares {
		err := middleware(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

var authMid = func(w http.ResponseWriter, r *http.Request) error {
	if r.FormValue("email") == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("Unauthorized")
	}
	return nil
}

var logMid = func(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Request received at ", r.URL.Path)
	return nil
}
