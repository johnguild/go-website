package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/johnguild/go-website/internal/tokenator"
)

func GetLogoutHandler(w http.ResponseWriter, r *http.Request) {
	newCookie, _ := tokenator.GenerateCookieWithToken(nil)
	http.SetCookie(w, newCookie)

	newUrl := fmt.Sprintf("%s://%s", strings.Split(r.Proto, "/")[0], r.Host)
	// fmt.Println(newUrl)
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", newUrl)
	} else {
		w.Header().Set("Location", newUrl)
	}
	w.WriteHeader(http.StatusOK)
}
