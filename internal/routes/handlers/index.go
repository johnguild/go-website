package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// homeTemplate.Execute(w, "ws://"+r.Host+"/server-log")
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}
