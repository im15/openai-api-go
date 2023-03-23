package openai

import (
	"context"
	"testing"
)

func TestClient_CreateEdit(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.CreateEdit(context.Background(), EditRequestBody{
		Model:       TextDavinciEdit001,
		Instruction: "Fix the spelling mistakes",
		Input:       "What day of the wek is it?",
	})
	if err != nil {
		t.Fatalf("create edit error: %v", err)
	}
	t.Logf("%v", body)
}
