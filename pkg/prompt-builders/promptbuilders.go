package promptbuilders

import "fmt"

type PromptBuilder struct {
}

func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{}
}

func (builder PromptBuilder) Build(input string, replyOption string) string {
	switch replyOption {
	case JokeReplyOption:
		return buildJoke(input)
	case SimpleReplyOption:
		return buildSimpleReply(input)
	case TwoSentenceHorrorRU:
		return buildForTwoSentenceHorrorStoryRU(input)
	default:
		return buildSimpleReply(input)
	}
}

func buildJoke(nakedPrompt string) string {
	prePrompt := "Generate funny joke about topic: "
	return prePrompt + nakedPrompt
}

func buildSimpleReply(nakedPrompt string) string {
	return nakedPrompt
}

func buildForTwoSentenceHorrorStoryRU(theme string) string {
	return fmt.Sprintf(`Тема: Завтрак
						История ужасов в двух предложениях: Он всегда перестает плакать, когда я наливаю молоко в его хлопья. Я просто должна помнить, чтобы он не видел своего лица на коробке.
							
						Тема: %s
						История ужасов в двух предложениях:`, theme)
}
