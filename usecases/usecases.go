package usecases

import (
	"github.com/sashabaranov/go-openai"
)

type GenerateTextUseCase struct {
	generator     Generator
	promptBuilder PrompBuilder
}

type ContextMessage struct {
	IsFromUser bool   `json:"isFromUser"`
	Message    string `json:"message"`
}

type ExecutionResult struct {
	Content          string `json:"content"`
	Model            string `json:"model"`
	TokenUsage       int    `json:"tokenUsage"`
	TokenUsageInput  int    `json:"tokenUsageInput"`
	TokenUsageOutput int    `json:"tokenUsageOutput"`
}

func NewGenerateTextUseCase(generator Generator, promptBuilder PrompBuilder) *GenerateTextUseCase {
	return &GenerateTextUseCase{
		generator:     generator,
		promptBuilder: promptBuilder,
	}
}

func (uc *GenerateTextUseCase) Execute(input string, replyOption string, model string, context []ContextMessage) (ExecutionResult, error) {
	prompt := uc.promptBuilder.Build(input, replyOption)

	var messagesConverted []openai.ChatCompletionMessage

	if context != nil {
		messagesConverted = convertMessagesToChatCompletionMessages(context)
	}

	chatCompResult, err := uc.generator.Generate(prompt, model, messagesConverted)
	if err != nil {
		return ExecutionResult{}, err
	}

	result := ExecutionResult{
		Content:          chatCompResult.Choices[0].Message.Content,
		Model:            chatCompResult.Model,
		TokenUsage:       chatCompResult.Usage.TotalTokens,
		TokenUsageInput:  chatCompResult.Usage.PromptTokens,
		TokenUsageOutput: chatCompResult.Usage.CompletionTokens,
	}

	return result, nil
}
