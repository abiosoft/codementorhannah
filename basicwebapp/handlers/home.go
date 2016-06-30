package handlers

import (
	"github.com/abiosoft/codementorhannah/basicwebapp/sessions"
	"html/template"
	"net/http"
)

// Handle is the root handler.
func Handle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		handleErr(w, r, err)
		return
	}

	values := map[string]interface{}{
		"user":     "Hannah",
		"loggedIn": sessions.LoggedIn(r),
	}

	if sessions.LoggedIn(r) {
		values["url"] = "/login"
	} else {
		values["url"] = "/logout"
	}

	err = t.Execute(w, values)
	if err != nil {
		handleErr(w, r, err)
	}

}
