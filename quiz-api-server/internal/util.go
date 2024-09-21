package internal

import (
    "encoding/json"
    "io"
    "os"
    "quiz-api-server/models"
)

func LoadQuestions(filename string) ([]models.Question, error) {
    jsonFile, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer jsonFile.Close()

    byteValue, _ := io.ReadAll(jsonFile)
    var questions []models.Question
    json.Unmarshal(byteValue, &questions)

    return questions, nil
}
