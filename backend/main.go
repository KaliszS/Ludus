package main

import (
	"log"
	"net/http"

	"ludus/middleware"
	"ludus/routes"
)


func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is working!"))
}

func main() {
	api := http.NewServeMux()
	userHandler := &routes.UserHandler{}

	api.HandleFunc("GET /", defaultHandler)
	api.HandleFunc("GET /user", userHandler.GetUser)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(api),
	}

	log.Println("Server is running on port 8080")
	server.ListenAndServe()
}