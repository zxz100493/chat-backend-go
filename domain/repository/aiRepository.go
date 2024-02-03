package repository

type AiRepository interface {
	Chat() string
	Response() string
	New(msg string) AiRepository
}
