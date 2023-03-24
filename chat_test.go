package openai

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func TestClient_CreateChatCompletion(t *testing.T) {
	c := NewClient(apiToken)
	body, err := c.CreateChatCompletion(context.Background(), ChatRequestBody{
		Model:  GPT35Turbo,
		Stream: true,
		Messages: []*ChatMessage{
			{Role: RoleUser, Content: "golang build -ldflags usage?"},
		},
	})
	if err != nil {
		t.Fatalf("Create chat completion error: %v", err)
	}
	t.Logf("%v", body)

	go func() {
		for {
			select {
			case chunk, ok := <-body.StreamChan:
				if !ok {
					return
				}
				//chunk.ID
				//chunk.Created
				//chunk.Object
				//chunk.Choices
				if len(chunk.Choices) > 0 {
					choice := chunk.Choices[0]
					if choice.FinishReason != nil {
						fmt.Print("\n\n")
						break
					}
					if choice.Message != nil {
						fmt.Printf("chunk: %v", choice.Message)
					} else if choice.Delta != nil {
						if choice.Delta.Role != "" {
							fmt.Printf("%s\n", choice.Delta.Role)
						} else if choice.Delta.Content != "" {
							fmt.Print(choice.Delta.Content)
						}
					}
				}

				//default:

			}
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	t.Logf("quit: %v", <-quit)
}
