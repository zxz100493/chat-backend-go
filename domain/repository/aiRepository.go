package repository

type AiRepository interface {
	Chat(msg string) string
}
