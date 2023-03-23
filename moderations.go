package openai

import (
	"context"
	"net/http"
)

// Moderations
// Given a input text, outputs if the model classifies it as violating OpenAI's
// content policy.

type ModerationRequestBody struct {
	Input string
	Model string `json:"model,omitempty"`
}

type ModerationObject struct {
	Hate            bool `json:"hate"`
	HateThreatening bool `json:"hate/threatening"`
	SelfHarm        bool `json:"self-harm"`
	Sexual          bool `json:"sexual"`
	SexualMinors    bool `json:"sexual/minors"`
	Violence        bool `json:"violence"`
	ViolenceGraphic bool `json:"violence/graphic"`
}

type ModerationResult struct {
	Categories     ModerationObject `json:"categories"`
	CategoryScores ModerationObject `json:"category_scores"`
	Flagged        bool             `json:"flagged"`
}

type ModerationResponseBody struct {
	ID      string             `json:"id"`
	Model   string             `json:"model"`
	Results []ModerationResult `json:"results"`
}

// CreateModeration Classifies if text violates OpenAI's Content Policy
// POST https://api.openai.com/v1/moderations
func (c *Client) CreateModeration(
	ctx context.Context,
	reqBody ModerationRequestBody) (resBody ModerationResponseBody, err error) {
	const apiURL = apiURLPrefix + "/v1/moderations"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}
