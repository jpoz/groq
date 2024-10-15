package main

import (
	"os"

	"github.com/jpenna/groq"
)

func main() {
	client := groq.NewClient()
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	response, err := client.CreateTranscription(groq.TranscriptionCreateParams{
		File: file,
	})
	if err != nil {
		panic(err)
	}

	println(response)
}
