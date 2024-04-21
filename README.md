

This Project provides a starting point for building a fullstack website using [Go](https://go.dev/) as the primary language with [HTMX](https://htmx.org/) and [Tailwindcss](https://tailwindcss.com/) for the frontend.


Features:

1. Public and Private Pages/Content with Authentication
    - Pages for guests like homepage, loginpage etc.. should be available to all
    - Authenticated pages/content will either redirect or return 404 not found

2. JWT Authentication.
    - Will be using [github.com/golang-jwt/jwt/v5](https://pkg.go.dev/github.com/golang-jwt/jwt/v5) for generating and validating token.
    - Token will be stored to session cookie named "token".

3. Middlewares Support
    - centralized route definition in routes.go 
    - Adding middlewares should be easy.
    ```
    mux.HandleFunc("/", handlers.IndexHandler)
    mux.HandleFunc("/dashboard", chainHandlers(
        middlewares.AuthenticatedContent,
        handlers.DashboardHandler,
    ))
    ```
4. Minimal Frontend Dependencies.
    - Focuses on server-side rendering with HTML responses (no JavaScript or JSON usage in this base template)


Limitations:
    
1. No .env or OS environment variable loading
2. No database implementation of any kind



Getting Started:

1. Download source code.
2. Open and go to source root directory.
3. Run ```go mod tidy```
4. Start server with ```go run ./cmd/website/```

