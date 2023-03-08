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
	model            = "text-davinci-003"
	temperature      = 0.6
	maxTokens        = 256
	topP             = 1
	frequencyPenalty = 0
	presencePenalty  = 0
	logprobs         = 0
)

type request struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float32 `json:"temperature"`
	MaxTokens        int8    `json:"max_tokens"`
	TopP             int8    `json:"top_p"`
	FrequencyPenalty int8    `json:"frequency_penalty"`
	PresencePenalty  int8    `json:"presence_penalty"`
	Logprobs         int8    `json:"logprobs"`
}

// type probs struct {
// 	Tokens        []string  `json:"tokens"`
// 	TokenLogprobs []float64 `json:"token_logprobs"`
// 	top_logprobs  []       `json:"top_logprobs"`
// 	TextOffset    []int     `json:"text_offset"`
// }

type choice struct {
	Text  string `json:"text"`
	Index int8   `json:"index"`
	// Logprobs     probs  `json:"logprobs,omitempty"`
	FinishReason string `json:"finish_reason"`
}

type usage struct {
	PromptTokens     int8 `json:"prompt_tokens"`
	CompletionTokens int8 `json:"completion_tokens"`
	TotalTokens      int8 `json:"total_tokens"`
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
		Model:            model,
		Prompt:           question,
		Temperature:      float32(temperature),
		MaxTokens:        int8(maxTokens),
		TopP:             int8(topP),
		FrequencyPenalty: int8(frequencyPenalty),
		PresencePenalty:  int8(presencePenalty),
		Logprobs:         int8(logprobs),
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
		"https://api.openai.com/v1/completions",
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

	var response Response

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("Response parse error", err.Error())

		return nil
	}

	// Print the string
	fmt.Println(string(bodyBytes))

	return &response
}
