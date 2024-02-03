package baidu

import (
	"chat-go/domain/repository"
	"chat-go/infrastructurre/ai"
	"chat-go/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	APIKey     string
	APISecret  string
	actionName = "baidu api"
)

type Baidu struct {
	ai.ChatMsg
}

func (b *Baidu) New(msg string) repository.AiRepository {
	return &Baidu{ai.ChatMsg{Msg: msg, Resp: ""}}
}

func init() {
	APIKey = os.Getenv("BAIDU_CHAT_AI_API_KEY")
	APISecret = os.Getenv("BAIDU_CHAT_AI_SECRET_KEY")
}

func (b *Baidu) Chat() string {
	defer util.TraceAction(actionName + "request chat")()

	msg := b.Msg
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro?access_token=" + GetAccessToken()

	question := `{"messages":[{"role":"user","content":"` + msg + `"}]}`
	headers := map[string]string{"Content-Type": "application/json"}
	response, err := makeRequest(url, "POST", headers, strings.NewReader(question))

	if err != nil {
		fmt.Println(err)
	}

	return response
}

func (b Baidu) Response() string {
	return ""
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", APIKey, APISecret)
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	defer util.TraceAction(actionName + "get access token")()

	response, err := makeRequest(url, "POST", headers, strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	accessTokenObj := map[string]string{}

	err = json.Unmarshal([]byte(response), &accessTokenObj)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return accessTokenObj["access_token"]
}

func makeRequest(url, method string, headers map[string]string, body io.Reader) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
