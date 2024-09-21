package cmd

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "github.com/spf13/cobra"
)

var comparisonUsername string

var getComparisonCmd = &cobra.Command{
    Use:   "get-comparison",
    Short: "Compare your quiz result to other users",
    Run: func(cmd *cobra.Command, args []string) {
        if comparisonUsername == "" {
            fmt.Println("Username is required for comparison")
            return
        }

        // Send GET request with username as a query parameter
        url := fmt.Sprintf("http://localhost:8080/compare?username=%s", comparisonUsername)
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Error fetching comparison:", err)
            return
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response:", err)
            return
        }

        if resp.StatusCode != http.StatusOK {
            fmt.Printf("Server responded with status: %s\n", resp.Status)
            fmt.Printf("Response body: %s\n", string(body))
            return
        }

        // Decode the response
        var result map[string]string
        err = json.Unmarshal(body, &result)
        if err != nil {
            fmt.Println("Error decoding response:", err)
            return
        }

        // Display the comparison message
        fmt.Println(result["message"])
    },
}

func init() {
    rootCmd.AddCommand(getComparisonCmd)

    // Add a flag for username
    getComparisonCmd.Flags().StringVarP(&comparisonUsername, "username", "u", "", "Specify the username for comparison (required)")
}
