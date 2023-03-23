package openai

import (
	"context"
	"testing"
)

func TestClient_CreateImage(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.CreateImage(context.Background(), ImageRequestBody{
		Prompt: "A cute baby sea otter",
		N:      3,
		Size:   Size1024,
	})
	if err != nil {
		t.Fatalf("Create image error: %v", err)
	}

	t.Logf("%v", body.Created)
	for i, url := range body.Data {
		t.Logf("[%d] %s", i, url)
	}
}

func TestClient_CreateImageEdit(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.CreateImageEdit(context.Background(), ImageEditRequestBody{
		Image:  "",
		Prompt: "",
		N:      3,
		//Size:           "",
		//ResponseFormat: "",
		//User:           "",
	})

	if err != nil {
		t.Fatalf("Create image edit error: %v", err)
	}

	t.Logf("%v", body.Created)
	for i, url := range body.Data {
		t.Logf("[%d] %s", i, url)
	}
}

func TestClient_CreateImageVariation(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.CreateImageVariation(context.Background(), ImageVariationRequestBody{
		Image: "",
		N:     3,
	})

	if err != nil {
		t.Fatalf("Create image variation error: %v", err)
	}

	t.Logf("%v", body.Created)
	for i, url := range body.Data {
		t.Logf("[%d] %s", i, url)
	}
}
