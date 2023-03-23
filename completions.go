package openai

import (
	"context"
	"net/http"
)

type CompletionRequestBody struct {
	Model           string         `json:"model"`
	Prompt          string         `json:"prompt,omitempty"`
	Suffix          string         `json:"suffix,omitempty"`
	MaxTokens       int            `json:"max_tokens,omitempty"`
	Temperature     float32        `json:"temperature,omitempty"`
	TopP            float32        `json:"top_p,omitempty"`
	N               int            `json:"n,omitempty"`
	Stream          bool           `json:"stream,omitempty"`
	Logprobs        int            `json:"logprobs,omitempty"`
	Echo            bool           `json:"echo,omitempty"`
	Stop            string         `json:"stop,omitempty"`
	PresencePenalty float32        `json:"presence_penalty,omitempty"`
	BestOf          int            `json:"best_of,omitempty"`
	LogitBias       map[string]int `json:"logit_bias,omitempty"`
	User            string         `json:"user,omitempty"`
}

type CompletionChoice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     *int   `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type CompletionResponseBody struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int                `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   TokensUsage        `json:"usage"`
}

func (c *Client) CreateCompletions(
	ctx context.Context,
	reqBody CompletionRequestBody) (resBody CompletionResponseBody, err error) {
	const apiURL = "https://api.openai.com/v1/completions"

	var req *http.Request
	if req, err = c.newRequest(ctx, POST, apiURL, reqBody); err != nil {
		return
	}

	if err = c.getRequest(req, resBody); err != nil {
		return
	}
	return
}
