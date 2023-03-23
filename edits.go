package openai

import (
	"context"
	"errors"
	"net/http"
)

// Edits
// Given a prompt and an instruction, the model will return an edited
// version of the prompt.

type EditRequestBody struct {
	Model       string  `json:"model"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
}

type EditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type EditResponseBody struct {
	Usage   TokensUsage  `json:"usage"`
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []EditChoice `json:"choices"`
}

// CreateEdit Creates a new edit for the provided inout, and parameters.
// POST https://api.openai.com/v1/edits
func (c *Client) CreateEdit(
	ctx context.Context,
	reqBody EditRequestBody) (resBody EditResponseBody, err error) {

	switch reqBody.Model {
	case TextDavinciEdit001:
	default:
		err = errors.New("invalid `model`")
		return
	}

	if reqBody.Instruction == "" {
		err = errors.New("`instruction` not provided")
		return
	}

	const apiURL = apiURLPrefix + "/v1/edits"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}
