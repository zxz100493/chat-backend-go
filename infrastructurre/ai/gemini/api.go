package gemini

import (
	"chat-go/app"
	"chat-go/domain/repository"
	"chat-go/infrastructurre/ai"
	"chat-go/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	APIKey     string
	actionName = "Gemini api"
)

type Gemini struct {
	ai.ChatMsg
}

type Response struct {
	Result string `json:"result"`
}

func init() {
	APIKey = os.Getenv("GENEMI_API_KEY")
}

func (g *Gemini) New(msg string) repository.AiRepository {
	fmt.Println("get genemi api key", APIKey)
	fmt.Println("app config", app.Config.Genemi.APIKey)

	if APIKey == "" {
		APIKey = app.Config.Genemi.APIKey
	}

	return &Gemini{ai.ChatMsg{Msg: msg, Resp: ""}}
}

func (g Gemini) Chat() string {
	defer util.TraceAction(actionName + "request chat")()

	msg := g.Msg

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(APIKey))
	if err != nil {
		log.Println("get client error", err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")

	fmt.Println("msg", msg)

	resp, err := model.GenerateContent(ctx, genai.Text(msg))
	if err != nil {
		log.Println(err)
	}

	log.Printf("resp %v", resp)

	if resp == nil {
		log.Println("本次没有输出内容")
		return getJSONStr("让我想一想")
	}

	str := ""

	for _, candidate := range resp.Candidates {
		if len(candidate.Content.Parts) > 0 {
			str += fmt.Sprint(candidate.Content.Parts[0])
			fmt.Println(str)
		}
	}

	return getJSONStr(str)
}

func (g Gemini) Response() string {
	return ""
}

func getJSONStr(str string) string {
	response := Response{
		Result: str,
	}

	// 将 Response 结构体转换为 JSON 字符串
	jsonStr, err := json.Marshal(response)
	if err != nil {
		log.Println("JSON 格式化错误", err)
		return ""
	}

	return string(jsonStr)
}
