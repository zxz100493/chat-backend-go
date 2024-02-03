package xunfei

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

/**
 *  WebAPI 接口调用示例 接口文档（必看）：https://www.xfyun.cn/doc/spark/Web.html
 * 错误码链接：https://www.xfyun.cn/doc/spark/%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E.html（code返回错误码时必看）
 * @author iflytek
 */

var (
	hostURL   string
	appid     string
	apiSecret string
	apiKey    string
)

func init() {
	hostURL = "wss://aichat.xf-yun.com/v1/chat"
	appid = os.Getenv("XUNFEI_APPID")
	apiSecret = os.Getenv("XUNFEI_APISECRET")
	apiKey = os.Getenv("XUNFEI_APIKEY")
}

type Xunfei struct {
}

func (x Xunfei) Chat(msgs string) string {
	const handshakeTimeout = 5 * time.Second

	d := websocket.Dialer{
		HandshakeTimeout: handshakeTimeout,
	}

	// 握手并建立websocket 连接
	conn, resp, err := d.Dial(assembleAuthURL1(hostURL, apiKey, apiSecret), nil)

	const webSocketStatusSwitchingProtocols = 101

	if err != nil {
		panic(readResp(resp) + err.Error())
	} else if resp.StatusCode != webSocketStatusSwitchingProtocols {
		panic(readResp(resp) + err.Error())
	}

	defer resp.Body.Close() // Close the response body

	go func() {
		data := genParams1(appid, msgs)
		err := conn.WriteJSON(data)

		if err != nil {
			fmt.Println("Error writing JSON:", err)
			return
		}
	}()

	var answer = ""
	// 获取返回的数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}

		var data map[string]interface{}
		err1 := json.Unmarshal(msg, &data)

		if err1 != nil {
			fmt.Println("Error parsing JSON:", err1)
			return ""
		}

		fmt.Println(string(msg))

		// 解析数据
		payload, ok := data["payload"].(map[string]interface{})
		if !ok {
			fmt.Println("Error parsing payload")
			return ""
		}

		choices, ok := payload["choices"].(map[string]interface{})
		if !ok {
			fmt.Println("Error parsing choices")
			return ""
		}

		header, ok := data["header"].(map[string]interface{})
		if !ok {
			fmt.Println("Error parsing header")
			return ""
		}

		code, ok := header["code"].(float64)
		if !ok {
			fmt.Println("Error parsing code")
			return ""
		}

		if code != 0 {
			fmt.Println(data["payload"])
			return ""
		}

		status, ok := choices["status"].(float64)
		if !ok {
			fmt.Println("Error parsing status")
			return ""
		}

		fmt.Println(status)

		text, ok := choices["text"].([]interface{})
		if !ok {
			fmt.Println("Error parsing text")
			return ""
		}

		content, ok := text[0].(map[string]interface{})["content"].(string)
		if !ok {
			fmt.Println("Error parsing text")
			return ""
		}

		const expectedStatus = 2

		if status != expectedStatus {
			answer += content
		} else {
			fmt.Println("收到最终结果")
			answer += content
			usage, ok := payload["usage"].(map[string]interface{})
			if !ok {
				fmt.Println("Error parsing usage")
				return ""
			}
			temp, ok := usage["text"].(map[string]interface{})
			if !ok {
				fmt.Println("Error parsing text")
				return ""
			}
			totalTokens, ok := temp["total_tokens"].(float64)
			if !ok {
				fmt.Println("Error parsing total_tokens")
				return ""
			}
			fmt.Println("total_tokens:", totalTokens)
			conn.Close()
			break
		}
	}

	// 输出返回结果
	fmt.Println(answer)

	time.Sleep(1 * time.Second)

	return answer
}

// 生成参数
func genParams1(appid, question string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名
	messages := []Message{
		{Role: "user", Content: question},
	}

	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      "general",    // 根据实际情况修改返回的数据结构和字段名
				"temperature": float64(0.8), // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(6),     // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(2048),  // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",    // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}

	return data // 根据实际情况修改返回的数据结构和字段名
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthURL1(hostURL, apiKey, apiSecret string) string {
	ul, err := url.Parse(hostURL)
	if err != nil {
		fmt.Println(err)
	}
	// 签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	// date = "Tue, 28 May 2019 09:10:42 MST"
	// 参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	// 拼接签名字符串
	sgin := strings.Join(signString, "\n")

	// 签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	// fmt.Println(sha)
	// 构建请求参数 此时不需要urlencoding
	authURL := fmt.Sprintf("hmac username=%q, algorithm=%q, headers=%q, signature=%q", apiKey,
		"hmac-sha256", "host date request-line", sha)
	// 将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authURL))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	// 将编码后的字符串url encode后添加到url后面
	callurl := hostURL + "?" + v.Encode()

	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
