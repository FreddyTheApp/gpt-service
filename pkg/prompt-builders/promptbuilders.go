package promptbuilders

import "fmt"

type PromptBuilder struct {
}

func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{}
}

func (builder PromptBuilder) BuildJoke(nakedPrompt string) string {
	prePrompt := "Generate funny joke about topic: "
	return prePrompt + nakedPrompt
}

func (b PromptBuilder) BuildSimpleReply(nakedPrompt string) string {
	return nakedPrompt
}

func (b PromptBuilder) BuildForTwoSentenceHorrorStoryRU(theme string) string {
	return fmt.Sprintf(`Тема: Завтрак
						История ужасов в двух предложениях: Он всегда перестает плакать, когда я наливаю молоко в его хлопья. Я просто должна помнить, чтобы он не видел своего лица на коробке.
							
						Тема: %s
						История ужасов в двух предложениях:`, theme)
}
