package openai

import (
	"context"
	"errors"
	"net/http"
)

// Audio
// Learn how to turn audio into text.

type AudioRequestBody struct {
	File           string  `json:"file"`
	Model          string  `json:"model"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float32 `json:"temperature,omitempty"`
	Language       string  `json:"language,omitempty"` // just create transcription
}

type AudioResponseBody struct {
	Text string `json:"text"`
}

// CreateTranscription Transcribes audio into the input language.
// POST https://api.openai.com/v1/audio/transcriptions
func (c *Client) CreateTranscription(
	ctx context.Context,
	reqBody AudioRequestBody) (resBody AudioResponseBody, err error) {
	if reqBody.File == "" {
		err = errors.New("`file` not provided")
		return
	}

	switch reqBody.Model {
	case Whisper1:
	default:
		err = errors.New("only `whisper-1` is currently available")
		return
	}

	const apiURL = apiURLPrefix + "/v1/audio/transcriptions"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}

// CreateTranslation Translates audio into English.
// POST https://api.openai.com/v1/audio/translations
func (c *Client) CreateTranslation(
	ctx context.Context,
	reqBody AudioRequestBody) (resBody AudioResponseBody, err error) {
	if reqBody.File == "" {
		err = errors.New("")
		return
	}

	switch reqBody.Model {
	case Whisper1:
	default:
		err = errors.New("only `whisper-1` is currently available")
		return
	}

	const apiURL = apiURLPrefix + "/v1/audio/translations"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}
