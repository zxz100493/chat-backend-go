package baidu

import (
	"encoding/json"
	"fmt"
)

func (b Baidu) ParseResult(jsonStr string) {
	var chatCompletion ChatCompletion
	err := json.Unmarshal([]byte(jsonStr), &chatCompletion)

	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}

	fmt.Println("baidu Ai:", chatCompletion.Result)
}
