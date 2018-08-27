package handler

import (
	"fmt"
	"net/http"
)

// FormHandler parses a form submitted over http(s)
// Returns "Hello <name>"
func FormHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte("Oops!"))
		return
	}
	name := req.Form.Get("name")
	if name == "" {
		name = "World"
	}
	res.Write([]byte(fmt.Sprintf("Hello %s", name)))
}
