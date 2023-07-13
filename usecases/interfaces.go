package usecases

import "github.com/sashabaranov/go-openai"

type Generator interface {
	Generate(string) (string, error)
	GenerateWithContext(string, []openai.ChatCompletionMessage) (string, error)
}

type PrompBuilder interface {
	BuildJoke(string) string
	BuildSimpleReply(string) string
	BuildForTwoSentenceHorrorStoryRU(string) string
}

type UseCase interface {
	Execute(string, ReplyOption, []Message) (string, error)
}
