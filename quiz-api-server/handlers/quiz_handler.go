package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quiz-api-server/internal"
	"quiz-api-server/models"
	"sync"
)

var quizResults = make(map[string]models.Result)
var resultMutex sync.Mutex

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := internal.LoadQuestions("data/questions.json")
	if err != nil {
		http.Error(w, "Error loading questions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Username string      `json:"username"`
		Answers  map[int]int `json:"answers"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	questions, err := internal.LoadQuestions("data/questions.json")
	if err != nil {
		http.Error(w, "Error loading questions", http.StatusInternalServerError)
		return
	}

	correctCount := 0
	totalQuestions := len(requestBody.Answers)

	for questionID, selectedAnswer := range requestBody.Answers {
		for _, q := range questions {
			if q.ID == questionID && q.CorrectAnswer == selectedAnswer {
				correctCount++
			}
		}
	}

	resultMutex.Lock()
	quizResults[requestBody.Username] = models.Result{
		UserName:       requestBody.Username,
		CorrectAnswers: correctCount,
		TotalQuestions: totalQuestions,
	}
	resultMutex.Unlock()

	response := fmt.Sprintf("You got %d out of %d correct!", correctCount, totalQuestions)
	json.NewEncoder(w).Encode(map[string]string{"message": response})
}

func GetComparison(w http.ResponseWriter, r *http.Request) {
	// Get the username from query params
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	resultMutex.Lock()
	defer resultMutex.Unlock()

	// Lookup the result for the given username in the map
	userResult, exists := quizResults[username]
	if !exists {
		http.Error(w, "No results found for the given username", http.StatusNotFound)
		return
	}

	// Calculate how many users the given user performed better than
	betterThanCount := 0
	totalUsers := len(quizResults)

	// Calculate the percentage of users this user outperformed
	var response string
	if totalUsers == 1 {
		response = "You are the only player so far"
	} else {
		for _, result := range quizResults {
			if userResult.CorrectAnswers > result.CorrectAnswers && userResult.UserName != result.UserName {
				betterThanCount++
			}
		}
		betterThanPercentage := float64(betterThanCount) / float64(totalUsers-1) * 100
		response = fmt.Sprintf("You were better than %.2f%% of all quizzers", betterThanPercentage)
	}

	// Send response back to the client
	json.NewEncoder(w).Encode(map[string]string{"message": response})
}
