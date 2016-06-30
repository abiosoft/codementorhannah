package main

import (
	"fmt"
	"net/http"

	"github.com/abiosoft/codementorhannah/basicwebapp/handlers"
)

func init() {
	registerHandlers()
}

func registerHandlers() {
	http.HandleFunc("/static/", handlers.HandleStatic)
	http.HandleFunc("/login", handlers.LoginUser)
	http.HandleFunc("/logout", handlers.LogoutUser)
	http.HandleFunc("/", handlers.Handle)
}

func main() {
	fmt.Println("Listening on 8081")
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
