package openai

import (
	"context"
	"os"
	"testing"
)

var apiToken string

func init() {
	apiToken = os.Getenv("OPENAI_API_KEY")
}

func TestClient_RetrieveModel(t *testing.T) {
	c := NewClient(apiToken)
	model, err := c.RetrieveModel(context.Background(), GPT35Turbo0310)
	if err != nil {
		t.Fatalf("Retrieve model error: %v", err)
	}
	t.Logf("%v", model)
}

func TestClient_ListModels(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.ListModels(context.Background())
	if err != nil {
		t.Fatalf("List models error: %v", err)
	}
	t.Logf("%v", body)
}
