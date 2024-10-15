package main

import (
	"github.com/jpenna/groq"
)

func main() {
	client := groq.NewClient() // will load API key from GROQ_API_KEY

	response, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-8b-8192",
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: "What is the meaning of life?",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	println(response.Choices[0].Message.Content)
}
