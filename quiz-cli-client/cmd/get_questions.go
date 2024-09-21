package cmd

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "quiz-cli-client/models"

    "github.com/spf13/cobra"
)

var getQuestionsCmd = &cobra.Command{
    Use:   "get-questions",
    Short: "Fetch quiz questions",
    Run: func(cmd *cobra.Command, args []string) {
        resp, err := http.Get("http://localhost:8080/questions")
        if err != nil {
            fmt.Println("Error fetching questions:", err)
            return
        }
        defer resp.Body.Close()

        body, _ := io.ReadAll(resp.Body)
        var questions []models.Question
        err = json.Unmarshal(body, &questions)
        if err != nil {
            fmt.Println("Error decoding response:", err)
            return
        }

        fmt.Println("Quiz Questions:")
        for _, q := range questions {
            fmt.Printf("Q%d: %s\n", q.ID, q.Question)
            for i, option := range q.Options {
                fmt.Printf("   %d) %s\n", i, option)
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(getQuestionsCmd)
}
