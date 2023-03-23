package openai

import (
	"context"
	"testing"
)

func TestClient_CreateChatCompletion(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.CreateChatCompletion(context.Background(), ChatRequestBody{
		Model: GPT35Turbo,
		Messages: []ChatMessage{
			{Role: RoleUser, Content: "golang build -ldflags usage?"},
		},
		User: "",
	})
	if err != nil {
		t.Fatalf("Create chat completion error: %v", err)
	}

	t.Logf("%v", body)
}
