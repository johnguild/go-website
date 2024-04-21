package handlers

import (
	"html/template"
	"net/http"

	"github.com/johnguild/go-website/internal/tokenator"
	"github.com/johnguild/go-website/internal/user"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Load userId from token
	cookie, _ := r.Cookie("token")
	userId := tokenator.GetCookieTokenClaimsValue(cookie.Value, "userId")

	// fetch user data if needed
	var tmp = &user.Credentials{Email: "", Password: "", Fullname: userId}

	t, _ := template.ParseFiles("templates/dashboard.html")
	t.Execute(w, tmp)
}
