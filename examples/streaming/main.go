package main

import (
	"fmt"

	"github.com/jpenna/groq"
)

func main() {
	client := groq.NewClient()

	chatCompletion, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-70b-8192",
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: "What is the meaning of life?",
			},
		},
		Stream:      true,
		Temperature: 1.0,
		MaxTokens:   1024,
		TopP:        1.0,
	})
	if err != nil {
		panic(err)
	}

	for delta := range chatCompletion.Stream {
		fmt.Print(delta.Choices[0].Delta.Content)
	}

	fmt.Print("\n")
}
