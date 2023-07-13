package usecases

import "github.com/sashabaranov/go-openai"

type ReplyOption int

const (
	JokeReplyOption ReplyOption = iota
	SimpleReplyOption
	TwoSentenceHorrorRU
)

type GenerateTextUseCase struct {
	generator     Generator
	promptBuilder PrompBuilder
}

type Message struct {
	IsFromUser bool   `json:"is-from-user"`
	Message    string `json:"message"`
}

func NewGenerateTextUseCase(generator Generator, promptBuilder PrompBuilder) *GenerateTextUseCase {
	return &GenerateTextUseCase{
		generator:     generator,
		promptBuilder: promptBuilder,
	}
}

func (uc *GenerateTextUseCase) Execute(input string, replyOption ReplyOption, messages []Message) (string, error) {
	var prompt string

	switch replyOption {
	case JokeReplyOption:
		prompt = uc.promptBuilder.BuildJoke(input)
	case SimpleReplyOption:
		prompt = uc.promptBuilder.BuildSimpleReply(input)
	case TwoSentenceHorrorRU:
		prompt = uc.promptBuilder.BuildForTwoSentenceHorrorStoryRU(input)
	default:
		prompt = uc.promptBuilder.BuildSimpleReply(input)
	}

	if messages != nil {
		messagesConverted := convertMessagesToChatCompletionMessages(messages)
		return uc.generator.GenerateWithContext(prompt, messagesConverted)
	}

	return uc.generator.Generate(prompt)
}

func convertMessagesToChatCompletionMessages(messages []Message) []openai.ChatCompletionMessage {
	result := []openai.ChatCompletionMessage{}

	for _, m := range messages {
		var role string
		if m.IsFromUser {
			role = openai.ChatMessageRoleUser
		} else {
			role = openai.ChatMessageRoleAssistant
		}

		result = append(result, openai.ChatCompletionMessage{
			Role:    role,
			Content: m.Message,
		})
	}

	return result
}
