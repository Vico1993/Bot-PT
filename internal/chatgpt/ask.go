package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	model            = "gpt-3.5-turbo-0301"
	temperature      = 1
	maxTokens        = 256
	topP             = 1
	frequencyPenalty = 0
	presencePenalty  = 0
	logprobs         = 0
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

//TopP             int8       `json:"top_p"`
//FrequencyPenalty int8       `json:"frequency_penalty"`
// PresencePenalty  int8       `json:"presence_penalty"`
// Logprobs         int8       `json:"logprobs"`

// type probs struct {
// 	Tokens        []string  `json:"tokens"`
// 	TokenLogprobs []float64 `json:"token_logprobs"`
// 	top_logprobs  []       `json:"top_logprobs"`
// 	TextOffset    []int     `json:"text_offset"`
// }

type choice struct {
	Message      Message `json:"message"`
	Index        int8    `json:"index"`
	FinishReason string  `json:"finish_reason"`
}

type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Response struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []choice `json:"choices,omitempty"`
	Usage   usage    `json:"usage"`
}

func Ask(question string) *Response {
	// Marshal the user object into a JSON-encoded byte slice
	body, err := json.Marshal(request{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: question,
			},
		},
		Temperature: float32(temperature),
		MaxTokens:   1000,
	})

	if err != nil {
		fmt.Println("Request parse error")
		return nil
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request with the desired URL and HTTP method
	req, err := http.NewRequest(
		"POST",
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_TOKEN"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed")
		return nil
	}
	defer resp.Body.Close()

	// Read the response body into a byte slice
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the string
	fmt.Println(string(bodyBytes))

	var response Response

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("Response parse error", err.Error())

		return nil
	}

	return &response
}
