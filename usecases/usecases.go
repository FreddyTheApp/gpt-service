package usecases

type ReplyOption int

const (
	JokeReplyOption ReplyOption = iota
	SimpleReplyOption
)

type GenerateTextUseCase struct {
	generator     Generator
	promptBuilder PrompBuilder
}

func NewGenerateTextUseCase(generator Generator, promptBuilder PrompBuilder) *GenerateTextUseCase {
	return &GenerateTextUseCase{
		generator:     generator,
		promptBuilder: promptBuilder,
	}
}

func (uc *GenerateTextUseCase) Execute(input string, replyOption ReplyOption) (string, error) {
	var prompt string

	switch replyOption {
	case JokeReplyOption:
		prompt = uc.promptBuilder.BuildJoke(input)
	case SimpleReplyOption:
		prompt = uc.promptBuilder.BuildSimpleReply(input)
	default:
		prompt = uc.promptBuilder.BuildSimpleReply(input)
	}

	return uc.generator.Generate(prompt)
}
