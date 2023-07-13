package promptbuilders

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
