package usecases

import (
	"github.com/sashabaranov/go-openai"
)

type Generator interface {
	Generate(input, model string, context []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error)
}

type PrompBuilder interface {
	Build(input, replyOption string) string
}

type UseCase interface {
	Execute(input, replyOption, model string, context []ContextMessage) (ExecutionResult, error)
}
