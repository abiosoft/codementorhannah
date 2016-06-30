package basictesting

import (
	"errors"
	"strconv"
)

func sayHello(name string) string {
	return "Hello " + name
}

func sayHi(name string) string {
	return "Hi " + name
}

func sayWhatsUp(name string) string {
	return "What's up" + name
}

func convertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

type Hello struct {
	name string
}

func (h *Hello) Say() string {
	return h.name
}

func NewHello() (*Hello, error) {
	return &Hello{}, MyError("My Error")
}

var YourError = errors.New("Your Error")

type MyError string

func (m MyError) Error() string {
	return "This is my error " + string(m)
}
