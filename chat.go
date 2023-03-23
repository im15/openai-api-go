package openai

import "context"

const (
	RoleUser      = "user"
	RoleSystem    = "system"
	RoleAssistant = "assistant"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequestBody struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	//LogitBias
	User string `json:"user,omitempty"`
}

type ChatChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type ChatResponseBody struct {
	Usage   TokensUsage  `json:"usage"`
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []ChatChoice `json:"choices"`
}

// CreateChatCompletion Create a completion for the chat message
func (c *Client) CreateChatCompletion(ctx context.Context, body ChatRequestBody) (*ChatResponseBody, error) {
	//const apiUrlV1 = "https://api.openai.com/v1/chat/completions"
	const apiUrlV1 = "http://localhost:8080/v1/chat/completions"
	switch body.Model {
	case GPT4, GPT40314, GPT432k, GPT432k0314, GPT35Turbo, GPT35Turbo0310:
	default:
		return nil, ErrInvalidModel
	}

	req, err := c.newRequest(ctx, POST, apiUrlV1, body)
	if err != nil {
		return nil, err
	}

	var resBody ChatResponseBody
	if err := c.getRequest(req, &resBody); err != nil {
		return nil, err
	}
	return &resBody, nil
}
