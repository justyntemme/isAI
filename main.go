package main

import (
	"fmt"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// APIResponse struct to hold response to send to user
type APIResponse struct {
	Likelihood float64 `json:"likelihood"`
}

func main() {
	var cmd = &cobra.Command{
		Use:   "ai-likelihood",
		Short: "Calculate the likelihood that a given text was written by an AI",
		Run:   calculateLikelihood,
	}

	// add the text and key flags to the command
	cmd.Flags().StringP("text", "t", "", "the text to evaluate")
	cmd.Flags().StringP("key", "k", "", "the OpenAI API key")

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func calculateLikelihood(cmd *cobra.Command, args []string) {
	// get the API key from the command parameter or the environment variable
	apiKey, _ := cmd.Flags().GetString("key")
	if apiKey == "" {
		apiKey = os.Getenv("API_KEY")
		if apiKey == "" {
			fmt.Println("API key not set")
			os.Exit(1)
		}
	}

	// get the text argument from the command
	text, err := cmd.Flags().GetString("text")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	client := gpt3.NewClient(apiKey)

	// call the OpenAI API to determine the likelihood that the text was written by an AI

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt: []string{"Was this paragraph written by an AI? '" + text + "'"},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(resp.Choices[0].Text)
}
