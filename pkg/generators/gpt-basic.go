package generators

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

const (
	GPT4  = "gpt-4"
	GPT35 = "gpt-3.5-turbo"
)

type GPTGenerator struct {
	token  string
	client *openai.Client
	model  string
}

func NewGPTBasicGenerator(token string, model string) *GPTGenerator {
	return &GPTGenerator{
		token:  token,
		client: openai.NewClient(token),
		model:  model,
	}
}

func (g *GPTGenerator) Generate(prompt string) (string, error) {
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
	return createCompletion(message, g.client, g.model)
}

func (g *GPTGenerator) GenerateWithContext(prompt string, context []openai.ChatCompletionMessage) (string, error) {
	messages := append(context, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	return createCompletion(messages, g.client, g.model)
}

func createCompletion(messages []openai.ChatCompletionMessage, client *openai.Client, model string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	if err != nil {
		return fmt.Sprintf("ChatCompletion error: %v\n", err), err
	}

	return resp.Choices[0].Message.Content, nil
}
