package routes

import (
	"net/http"

	"github.com/johnguild/go-website/internal/middlewares"
	"github.com/johnguild/go-website/internal/routes/handlers"
)

/*
*

	Self imposed Route rules
	1 route = 1 file
	Filename should be method+Action
		ex. postLogin, getUsers, putUser
	Each file should have 2 functions and 1 type struct
	1 type struct corresponds to request body validation and errors
	1 Handler function for ex. PostLoginHandler, GetLogoutHandler
		check method and other header stuff
		validate session cookie token (if private content)
		validate request body
		do your logic
		regenerate jwt token and attach to session cookie
		always return html tag or redirect
	2 Body function to build the type and validate input
*
*/

func AttachRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/dashboard", chainHandlers(
		middlewares.AuthenticatedContent,
		handlers.DashboardHandler,
	))
	mux.HandleFunc("/private", chainHandlers(
		middlewares.AuthenticatedContent,
		handlers.PrivateHandler,
	))
	mux.HandleFunc("/login", chainHandlers(
		middlewares.FormPostOnly,
		handlers.PostLoginHandler,
	))
	mux.HandleFunc("/logout", chainHandlers(
		middlewares.AuthenticatedContent,
		handlers.GetLogoutHandler,
	))
}

/*
This function takes multiple handler functions as arguments and returns a new
function that chains their execution.
*/
func chainHandlers(handlers ...func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			// Stop next function if a response status code has been set
			// LIMITATION: cant check current writer status but
			// having content-type set means the status is set.
			if status := w.Header().Get("Content-Type"); status != "" {
				return
			}

			handler(w, r)
		}
	}
}
