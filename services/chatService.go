package services

import (
	"chat-go/domain/repository"
	"chat-go/infrastructurre/ai/gemini"
	"chat-go/infrastructurre/ai/xunfei"
)

func ChatWithAi(msg string) interface{} {
	aiMode := "baidu"
	var aiSvc repository.AiRepository

	if aiMode == "baidu" {
		aiSvc = gemini.Gemini{}
		// aiSvc = baidu.Baidu{}
	} else if aiMode == "xunfei" {
		aiSvc = xunfei.Xunfei{}
	} else if aiMode == "gemini" {
		aiSvc = gemini.Gemini{}
	}

	return aiSvc.Chat(msg)
}
