package services

import (
	"chat-go/domain/repository"
	"chat-go/infrastructurre/ai/gemini"
	"chat-go/infrastructurre/ai/xunfei"
)

var AiServiceMap = map[string]repository.AiRepository{
	"baidu":  gemini.Gemini{},
	"xunfei": xunfei.Xunfei{},
	"gemini": gemini.Gemini{},
}

func ChatWithAi(msg string) interface{} {
	aiMode := "baidu"
	aiSvc, exists := AiServiceMap[aiMode]
	if !exists {
		aiSvc = AiServiceMap["gemini"]
	}
	return aiSvc.Chat(msg)
}
