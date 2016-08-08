package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abiosoft/codementorhannah/basicwebapp/model"
	"github.com/abiosoft/codementorhannah/basicwebapp/sessions"
)

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "user" && password == "pass" {

		err := sessions.SetValue(r, w, "authenticated", true)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func DumpUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, users)

}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	sessions.Clear(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
