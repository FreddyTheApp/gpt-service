package main

import (
	"os"

	"github.com/FreddyTheApp/gpt-service/pkg/generators"
	promptbuilders "github.com/FreddyTheApp/gpt-service/pkg/prompt-builders"
	"github.com/FreddyTheApp/gpt-service/usecases"
)

func main() {
	gptGenerator := generators.NewGPTBasicGenerator(os.Getenv("OPENAI_TOKEN"))
	basicEmailPromptBuilder := promptbuilders.NewPromptBuilder()
	basicUseCase := usecases.NewGenerateTextUseCase(gptGenerator, basicEmailPromptBuilder)

	app := NewApp(basicUseCase)
	app.Start()
}
