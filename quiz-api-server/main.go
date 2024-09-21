package main

import (
    "log"
    "net/http"
    "quiz-api-server/handlers"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/questions", handlers.GetQuestions).Methods("GET")
    r.HandleFunc("/submit", handlers.SubmitAnswers).Methods("POST")
    r.HandleFunc("/compare", handlers.GetComparison).Methods("GET")

    // Start the server
    log.Println("Quiz API server is running on port :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
