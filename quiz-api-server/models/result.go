package models

type Result struct {
    UserName      string `json:"user_name"`
    CorrectAnswers int    `json:"correct_answers"`
    TotalQuestions int    `json:"total_questions"`
}
