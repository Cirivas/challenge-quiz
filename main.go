package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	headers := handlers.AllowedHeaders([]string{"Contety-Type", "Access-Control-Allow-Headers", "Authorization"})
	methods := handlers.AllowedMethods([]string{"DELETE", "POST", "GET", "OPTIONS", "PUT", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	loggedRouter := handlers.LoggingHandler(os.Stdout, mux.NewRouter())

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(loggedRouter)))
}
