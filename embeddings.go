package openai

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

func (c *Client) CreateEmbeddings() {
	// POST https://api.openai.com/v1/embeddings
}
