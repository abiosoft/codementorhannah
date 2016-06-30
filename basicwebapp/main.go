package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

const sessionName = "OUR_SESSION"

func main() {

	// http.Handle("/static", http.FileServer(http.Dir("./")))

	http.HandleFunc("/static/", handleStatic)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/", handle)

	fmt.Println("Listening on 8081")

	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}

func handle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		handleErr(w, r, err)
		return
	}

	values := map[string]interface{}{
		"user":     "Hannah",
		"loggedIn": loggedIn(r),
	}

	if loggedIn(r) {
		values["url"] = "/login"
	} else {
		values["url"] = "/logout"
	}

	err = t.Execute(w, values)
	if err != nil {
		handleErr(w, r, err)
	}

}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	http.ServeFile(w, r, filename)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// session is retrieved from request
	session, err := store.Get(r, sessionName)

	if err != nil {
		handleErr(w, r, err)
		return
	}

	if username == "user" && password == "pass" {
		session.Values["authenticated"] = true
		err := session.Save(r, w)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func handleErr(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func loggedIn(r *http.Request) bool {
	session, _ := store.Get(r, sessionName)
	return session.Values["authenticated"] == true
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	for key := range session.Values {
		delete(session.Values, key)
	}
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
