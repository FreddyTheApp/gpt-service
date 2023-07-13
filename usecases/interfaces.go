package usecases

type Generator interface {
	Generate(string) (string, error)
}

type PrompBuilder interface {
	BuildJoke(string) string
	BuildSimpleReply(string) string
}

type UseCase interface {
	Execute(string, ReplyOption) (string, error)
}
