package openai

import (
	"context"
	"net/http"
)

type EmbeddingsRequestBody struct {
	Model string `json:"model"`
	Input string `json:"input"`
	User  string `json:"user,omitempty"`
}

type EmbeddingsResponseBody struct {
	Object string      `json:"object"`
	Model  string      `json:"model"`
	Usage  TokensUsage `json:"usage"`
	Data   []struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

// CreateEmbeddings
// POST https://api.openai.com/v1/embeddings
func (c *Client) CreateEmbeddings(
	ctx context.Context,
	reqBody EmbeddingsRequestBody) (resBody EmbeddingsResponseBody, err error) {
	const apiURL = apiURLPrefix + "/v1/embeddings"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}
	err = c.getRequest(req, &resBody)
	return
}
