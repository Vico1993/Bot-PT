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
	model       = "gpt-3.5-turbo-0301"
	temperature = 1
	// maxTokens        = 256
	// topP             = 1
	// frequencyPenalty = 0
	// presencePenalty  = 0
	// logprobs         = 0
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

type Choice struct {
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
	Choices []Choice `json:"choices,omitempty"`
	Usage   usage    `json:"usage"`
}

// Marshal the user object into a JSON-encoded byte slice
func buildReqBody(question string) []byte {
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

	return body
}

// Create a new request with the desired URL and HTTP method
func buildRequest(body []byte) *http.Request {
	req, err := http.NewRequest(
		"POST",
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(body),
	)

	if err != nil {
		fmt.Println("Error building request:", err.Error())
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_TOKEN"))

	return req
}

func Ask(question string) *Response {
	body := buildReqBody(question)
	if body == nil {
		return nil
	}

	req := buildRequest(body)

	// Create a new HTTP client
	client := &http.Client{}

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
