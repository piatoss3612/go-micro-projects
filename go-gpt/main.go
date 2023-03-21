package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API_KEY not found")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	rootCmd := &cobra.Command{
		Use:   "gpt3",
		Short: "GPT-3 CLI",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Say something(quit to end): ")
				if !scanner.Scan() {
					break
				}

				question := scanner.Text()

				switch question {
				case "quit":
					quit = true
				default:
					err := client.CompletionStreamWithEngine(
						ctx,
						gpt3.TextDavinci003Engine,
						gpt3.CompletionRequest{
							Prompt:      []string{question},
							MaxTokens:   gpt3.IntPtr(200),
							Temperature: gpt3.Float32Ptr(0.5),
						},
						func(cr *gpt3.CompletionResponse) {
							fmt.Print(cr.Choices[0].Text)
						})
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println()
				}
			}
		},
	}

	rootCmd.Execute()
}
