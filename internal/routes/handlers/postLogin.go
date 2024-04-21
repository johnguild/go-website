package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/johnguild/go-website/internal/tokenator"
	"github.com/johnguild/go-website/internal/user"
	"github.com/johnguild/go-website/internal/validators"
)

type LoginBody struct {
	Email       string `validate:"required,email,min=10,max=64"`
	EmailErr    string
	Password    string `validate:"min=8,max=32"`
	PasswordErr string
	MainErr     string
	HasError    bool
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Set 200 OK for successful handling,
	// regardless of 400/401
	// just return the proper html template
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Prepare template error
	templateError, _ := template.ParseFiles("templates/login/error.html")

	// Validate request body
	var body = FormBody(r)
	if body.HasError {
		templateError.Execute(w, body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if user credentials exists
	user, err := user.FindMatch(body.Email, body.Password)
	if err != nil {
		body.MainErr = "No user found with those credentials"
		templateError.Execute(w, body)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Generate and attach jwt token to session cookie token
	cookie, err := tokenator.GenerateCookieWithToken(user)
	if err != nil {
		fmt.Println(err.Error())
		body.MainErr = "Failed processing request!"
		templateError.Execute(w, body)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Set the cookie in the response
	http.SetCookie(w, cookie)

	newUrl := fmt.Sprintf("%s://%s/%s", strings.Split(r.Proto, "/")[0], r.Host, "dashboard")
	// fmt.Println(newUrl)
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", newUrl)
	} else {
		w.Header().Set("Location", newUrl)
	}
	w.WriteHeader(http.StatusOK)
}

func FormBody(r *http.Request) *LoginBody {
	var body = LoginBody{
		Email:       r.PostFormValue("email"),
		EmailErr:    "",
		Password:    r.PostFormValue("password"),
		PasswordErr: "",
		HasError:    false,
	}

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		body.HasError = true
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Email":
				body.EmailErr = validators.ErrorMessage(err.Field(), err.Tag(), err.Param())
			case "Password":
				body.PasswordErr = validators.ErrorMessage(err.Field(), err.Tag(), err.Param())
			}
		}
	}
	return &body
}
