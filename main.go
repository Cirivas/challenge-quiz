package main

import (
	"log"
	"net/http"
	"os"

	redis_db "github.com/cirivas/challenge-quiz/infrastructure/database/redis"
	"github.com/cirivas/challenge-quiz/registry"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	headers := handlers.AllowedHeaders([]string{"Contety-Type", "Access-Control-Allow-Headers", "Authorization"})
	methods := handlers.AllowedMethods([]string{"DELETE", "POST", "GET", "OPTIONS", "PUT", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	dbClient := redis_db.NewRedisClient()
	if err := dbClient.Connect(); err != nil {
		panic(err)
	}
	defer dbClient.Close()

	registry := registry.NewRegistry()
	controllers := registry.NewController(dbClient)

	router := mux.NewRouter()

	router.Handle("/quiz/{quizId}", http.HandlerFunc(controllers.QuizController.GetQuiz)).Methods("GET")
	router.Handle("/quiz/{quizId}/ranking/{respondent}", http.HandlerFunc(controllers.RankingController.GetRanking)).Methods("GET")

	router.Handle("/answer", http.HandlerFunc(controllers.QuizController.AnswerQuiz)).Methods("POST")

	router.Use(contentTypeApplicationJsonMiddleware)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(loggedRouter)))
}

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
