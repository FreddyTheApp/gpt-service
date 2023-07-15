package generators

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

const (
	GPT4  = "gpt-4"
	GPT35 = "gpt-3.5-turbo"
)

type GPTGenerator struct {
	token  string
	client *openai.Client
}

func NewGPTBasicGenerator(token string) *GPTGenerator {
	return &GPTGenerator{
		token:  token,
		client: openai.NewClient(token),
	}
}

func (g *GPTGenerator) Generate(prompt string, model string, context []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	messages := append(context, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	return createCompletion(messages, g.client, model)
}

func createCompletion(messages []openai.ChatCompletionMessage, client *openai.Client, model string) (openai.ChatCompletionResponse, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)

	return resp, err
}
