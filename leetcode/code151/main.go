package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// fmt.Println(reverseWords("a good   example"))

	client := openai.NewClient("REMOVED_OPENAI_KEY")

	req := openai.ChatCompletionRequest{
		Model: "gpt-4o", // Use "gpt-4.1" if specifically available
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Write a one-sentence bedtime story about a unicorn.",
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatalf("ChatCompletion error: %v", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func reverseWords(s string) string {

	splitted := strings.Split(s, " ")

	s = ""
	for i := len(splitted) - 1; i >= 0; i-- {
		if splitted[i] == "" {
			continue
		}
		s += splitted[i] + " "
	}

	s = strings.TrimSuffix(s, " ")

	return s
}
