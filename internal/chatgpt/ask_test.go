package chatgpt

import (
	"encoding/json"
	"testing"

	"github.com/Vico1993/Bot-PT/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestBuildReqBody(t *testing.T) {
	question := "What is the meaning of life?"
	expectedBody, _ := json.Marshal(request{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: question}},
		Temperature: float32(temperature),
		MaxTokens:   1000,
	})

	body := buildReqBody(question)

	assert.Equal(t, expectedBody, body, "Unexpected request body.")
}

func TestBuildRequest(t *testing.T) {
	type TestStruct struct {
		test bool
	}

	keyNeeed := []string{"Content-Type", "Authorization"}

	body, _ := json.Marshal(TestStruct{test: true})

	req := buildRequest(body)

	for key := range req.Header {
		assert.True(t, utils.InSlice(key, keyNeeed), "Key should be present in the header: "+key)
	}
}
