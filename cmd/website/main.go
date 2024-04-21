package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/johnguild/go-website/internal/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.AttachRoutes(mux)

	fmt.Println("server in http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
