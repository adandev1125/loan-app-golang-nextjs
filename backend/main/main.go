package main

import (
	"fmt"
	"main/main/config"

	"main/main/routes"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	fmt.Printf("Server started.\nYou can test on http://localhost:%d\n", config.PORT)

	router := routes.GetRouter()

	// cors middleware
	methods := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	handler := handlers.CORS(methods, origins, headers)(router)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.PORT),
		Handler: handler,
	}

	s.ListenAndServe()
}
