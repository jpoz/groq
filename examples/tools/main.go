package main

import (
	"encoding/json"
	"fmt"

	"github.com/jpenna/groq"
)

// This currently isn't working. It seems to be an issue on their side:
// https://discord.com/channels/1207099205563457597/1237401687270883361

func main() {
	client := groq.NewClient(groq.WithDebug(true))

	chatCompletion, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-70b-8192",
		Messages: []groq.Message{
			{
				Role:    "system",
				Content: "You are a function calling LLM that uses the data extracted from the get_weather function to answer questions about the weather. Include the city and the state in your response.",
			},
			{
				Role:    "user",
				Content: "What is weather in Boulder, CO",
			},
		},
		Tools: []groq.Tool{
			{
				Function: groq.ToolFunction{
					Name:        "get_weather",
					Description: "Returns the current weather in a city",
					Parameters: map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"city": map[string]interface{}{
								"type":        "string",
								"description": "The name of the city",
							},
							"state": map[string]interface{}{
								"type":        "string",
								"description": "The name of the state",
							},
						},
						"required": []string{"team_name"},
					},
				},
				Type: "function",
			},
		},
		ToolChoice: groq.ToolChoiceAuto,
	})
	if err != nil {
		panic(err)
	}

	bts, err := json.Marshal(chatCompletion)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bts))
}
