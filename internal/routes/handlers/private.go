package handlers

import (
	"html/template"
	"net/http"

	"github.com/johnguild/go-website/internal/user"
)

func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/private.html")

	var user = &user.Credentials{Fullname: "Private User Here"}
	t.Execute(w, user)
}
