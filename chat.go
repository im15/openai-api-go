package openai

import (
	"context"
	"net/http"
)

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
	Model            string         `json:"model"`
	Messages         []ChatMessage  `json:"messages"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Stop             string         `json:"stop,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
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
// POST https://api.openai.com/v1/chat/completions
func (c *Client) CreateChatCompletion(
	ctx context.Context,
	reqBody ChatRequestBody) (resBody ChatResponseBody, err error) {
	switch reqBody.Model {
	case GPT4, GPT40314, GPT432k, GPT432k0314, GPT35Turbo, GPT35Turbo0310:
	default:
		err = ErrInvalidModel
		return
	}

	const apiURL = apiURLPrefix + "/v1/chat/completions"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)
	return
}
