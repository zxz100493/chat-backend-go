package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Gemini struct {
}

type Response struct {
	Result string `json:"result"`
}

func (g Gemini) Chat(msg string) string {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GENEMI_API_KEY")))
	if err != nil {
		log.Println("get client error", err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	fmt.Println("msg", msg)
	// resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	resp, err := model.GenerateContent(ctx, genai.Text(msg))
	if err != nil {
		log.Println(err)
	}
	log.Printf("resp %v", resp)
	if resp == nil {
		log.Println("本次没有输出内容")
		return getJsonStr("让我想一想")
	}
	str := ""
	for _, candidate := range resp.Candidates {

		if len(candidate.Content.Parts) > 0 {
			str += fmt.Sprint(candidate.Content.Parts[0])
			fmt.Println()
		}
	}
	return getJsonStr(str)
}

func getJsonStr(str string) string {
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
