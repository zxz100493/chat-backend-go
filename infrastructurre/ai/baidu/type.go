package baidu

type ChatCompletion struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Result           string `json:"result"`
	FinishReason     string `json:"finish_reason"`
	Created          int64  `json:"created"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
	Usage            Usage  `json:"usage"`
}

type Usage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}
