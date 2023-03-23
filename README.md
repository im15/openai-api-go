# OpenAI API for Go

This library provider Go clients for OpenAI API.

Installation

```shell
go get github.com/im15/go-openai
```

## Example:

```go
package main

import (
	"context"
	openai "github.com/im15/openai-api-go"
	"log"
	"os"
)

func main() {
	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	body, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatRequestBody{
			Model: openai.GPT35Turbo,
			Message: []openai.ChatMessage{
				{
					Role: openai.RoleUser, 
					Content: "Hello!" 
				},
            },
        },
	)
	if err != nil {
        log.Printf("error: %v", err)
		return
	}

	log.Printf("%s", body.Choices[0].Message.Content)
}
```