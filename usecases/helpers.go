package usecases

import "github.com/sashabaranov/go-openai"

func convertMessagesToChatCompletionMessages(messages []ContextMessage) []openai.ChatCompletionMessage {
	result := make([]openai.ChatCompletionMessage, len(messages))

	for i, m := range messages {
		role := getOpenAIRole(m.IsFromUser)

		result[i] = openai.ChatCompletionMessage{
			Role:    role,
			Content: m.Message,
		}
	}
	return result
}

func getOpenAIRole(isUser bool) string {
	if isUser {
		return openai.ChatMessageRoleUser
	}
	return openai.ChatMessageRoleAssistant
}

// TODO: Calculations
var modelPriceMapping = map[string]float32{
	"gpt-4-0613":             0.03,
	"gpt-4-32k-0613":         0.06,
	"gpt-3.5-turbo-0613":     0.0015,
	"gpt-3.5-turbo-16k-0613": 0.003,
}
