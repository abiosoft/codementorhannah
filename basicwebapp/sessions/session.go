package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
)

const sessionName = "OUR_SESSION"

// SessionStore
var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	store.Options.MaxAge = 60 * 15
}

// GetSession gets the current user session.
func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}

// GetValue retrieves a value in the session.
func GetValue(r *http.Request, name string) interface{} {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil
	}
	return session.Values[name]
}

// SetValue sets a value in the session.
func SetValue(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values[key] = value
	return sessions.Save(r, w)
}

// Clear clears the session.
func Clear(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	for key := range session.Values {
		delete(session.Values, key)
	}
	session.Save(r, w)
}

// LoggedIn checks if user is logged in.
func LoggedIn(r *http.Request) bool {
	session, _ := store.Get(r, sessionName)
	return session.Values["authenticated"] == true
}

/*







interface explanations
***/
type User struct {
	name string
}

func (u User) Greet() {
	println(u.name)
}

type SuperUser struct {
	User
	title string
}

// func (s SuperUser) Greet() {
// 	println(s.title)
// 	s.u.Greet()
// }

func init() {
	Greet(User{})
}

type Greeter interface {
	Greet()
}

func Greet(g Greeter) {
	g.Greet()
}
