package middlewares

import (
	"net/http"
	"strings"

	"github.com/johnguild/go-website/internal/tokenator"
)

// Check session cookie if valid=
// if not, redirect to homepage
func AuthenticatedContent(w http.ResponseWriter, r *http.Request) {
	// Get token from cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		newCookie, _ := tokenator.GenerateCookieWithToken(nil)
		http.SetCookie(w, newCookie)
		http.Redirect(w, r, r.Host, http.StatusSeeOther)
		return
	}

	// validate token and retrieve claims
	_, err = tokenator.ValidateCookieToken(cookie.Value)
	if err != nil {
		newCookie, _ := tokenator.GenerateCookieWithToken(nil)
		http.SetCookie(w, newCookie)
		http.Redirect(w, r, r.Host, http.StatusSeeOther)
		return
	}
}

// check if request if from a form with Method POST
// if not return 404 Not Found
func FormPostOnly(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Header)
	if r.Method != "POST" || strings.Split(r.Header.Get("Content-Type"), ";")[0] != "multipart/form-data" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
}
