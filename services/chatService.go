package services

import (
	"chat-go/domain/repository"
	"chat-go/infrastructurre/ai/baidu"
	"chat-go/infrastructurre/ai/xunfei"
)

func ChatWithAi(msg string) interface{} {
	aiMode := "baidu"
	var aiSvc repository.AiRepository

	if aiMode == "baidu" {
		aiSvc = baidu.Baidu{}
	} else if aiMode == "xunfei" {
		aiSvc = xunfei.Xunfei{}
	}

	return aiSvc.Chat(msg)
}
