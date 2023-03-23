package openai

import (
	"context"
	"testing"
)

func TestClient_ListFiles(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.ListFiles(context.Background())
	if err != nil {
		t.Fatalf("list files error: %v", err)
	}
	t.Logf("%v", body)
	for i, fileObject := range body.Data {
		t.Logf("[%d] %v", i, fileObject)
	}
}
