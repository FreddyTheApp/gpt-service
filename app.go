package main

import (
	"os"

	"github.com/FreddyTheApp/gpt-service/handlers"
	"github.com/FreddyTheApp/gpt-service/usecases"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

type App struct {
	env     string
	useCase usecases.UseCase
}

func NewApp(useCase usecases.UseCase) *App {
	return &App{
		env:     os.Getenv("ENV"),
		useCase: useCase,
	}
}

func (app App) Start() {
	h := handlers.NewGinHandler(app.useCase)
	r := gin.New()

	r.POST("/message", h.HandleSimpleReplyRequest)
	r.POST("/joke", h.HandleJokeReplyRequest)
	r.POST("/horror", h.HandleTwoSentenceHorrorStoryRURequest)

	r.Run()
}
