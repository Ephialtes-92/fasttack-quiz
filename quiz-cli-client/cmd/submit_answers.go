package cmd

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    //"strings"

    "github.com/spf13/cobra"
)

var answerStr map[string]int // Use StringToIntVar to accept input as "1=2,2=3"
var answers = make(map[int]int)
var username string

var submitAnswersCmd = &cobra.Command{
    Use:   "submit-answers",
    Short: "Submit your answers to the quiz",
    Run: func(cmd *cobra.Command, args []string) {

        // Validate that a username is provided
        if username == "" {
            fmt.Println("Username is required")
            return
        }

        // Parse answerStr into answers (map[int]int)
        for key, value := range answerStr {
            questionID, err := strconv.Atoi(key)
            if err != nil {
                fmt.Println("Invalid question ID:", key)
                return
            }
            answers[questionID] = value
        }

        if len(answers) == 0 {
            fmt.Println("No answers provided")
            return
        }

        // Prepare request body
        submissionData := map[string]interface{}{
            "username": username,
            "answers":  answers,
        }

        jsonData, err := json.Marshal(submissionData)
        if err != nil {
            fmt.Println("Error preparing request data:", err)
            return
        }

        // Send POST request
        resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            fmt.Println("Error submitting answers:", err)
            return
        }
        defer resp.Body.Close()

        var result map[string]string
        json.NewDecoder(resp.Body).Decode(&result)

        fmt.Println(result["message"])
    },
}

func init() {
    rootCmd.AddCommand(submitAnswersCmd)

    // Use StringToIntVar to accept --answers flag as "1=2,2=3"
    submitAnswersCmd.Flags().StringToIntVar(&answerStr, "answers", nil, "Submit answers in the format --answers 1=2,2=3")
    submitAnswersCmd.Flags().StringVarP(&username, "username", "u", "", "Submit your username (required)")
}
