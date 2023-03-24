package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type Error struct {
	Code       *string `json:"code,omitempty"`
	Param      *string `json:"param,omitempty"`
	Message    string  `json:"message"`
	Type       string  `json:"type"`
	StatusCode int     `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

type RequestError struct {
	StatusCode int
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status code %d", r.StatusCode)
}

func (r *RequestError) Unwrap() error {
	return r
}

type ErrorResponseBody struct {
	Error *Error `json:"error,omitempty"`
}

var (
	ErrInvalidModel = errors.New("invalid model")
)

type TokensUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

const (
	apiURLPrefix = "https://api.openai.com"
)

type Client struct {
	token string
	OrgID string
}

func NewClient(token string) *Client {
	return NewClientWithOrg(token, "")
}

func NewClientWithOrg(token, orgID string) *Client {
	return &Client{
		token: token,
		OrgID: orgID,
	}
}

func (c *Client) newRequest(ctx context.Context,
	method string,
	url string,
	body any) (req *http.Request, err error) {
	var (
		data              io.Reader
		headerAccept      = "application/json; charset=utf-8"
		headerContentType = "application/json; charset=utf-8"
	)

	if body != nil {
		var buf bytes.Buffer
		if b, ok := body.(ImageEditRequestBody); ok {
			w := multipart.NewWriter(&buf)
			if err = b.WriteForm(w); err != nil {
				return
			}
			headerContentType = w.FormDataContentType()
		} else if b, ok := body.(ImageVariationRequestBody); ok {
			w := multipart.NewWriter(&buf)
			if err = b.WriteForm(w); err != nil {
				return
			}
			headerContentType = w.FormDataContentType()
		} else {
			if err = json.NewEncoder(&buf).Encode(body); err != nil {
				return
			}
		}
		data = &buf
	}

	if req, err = http.NewRequestWithContext(ctx, method, url, data); err != nil {
		return
	}
	if b, ok := body.(ChatRequestBody); ok {
		if b.Stream {
			headerAccept = "text/event-stream"
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Connection", "keep-alive")
		}
	}
	if c.OrgID != "" {
		req.Header.Set("OpenAI-Organization", c.OrgID)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Accept", headerAccept)
	req.Header.Set("Content-Type", headerContentType)

	return
}

func (c *Client) getRequest(req *http.Request, v any) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	var wasStreamResponse bool
	defer func() {
		if !wasStreamResponse {
			_ = res.Body.Close()
		}
	}()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("%v", err)
		} else {
			log.Printf("%s", body)
		}

		return errors.New(res.Status)
	}

	if v != nil {
		if b, ok := v.(*RetrieveFileContentResponseBody); ok {
			b.Data, err = io.ReadAll(res.Body)
			return err
		}
		if b, ok := v.(*ChatResponseBody); ok {
			if res.Header.Get("Content-Type") == "text/event-stream" {
				if b.StreamChan != nil {
					wasStreamResponse = true

					go func(b *ChatResponseBody, r io.ReadCloser) {
						defer func() {
							_ = r.Close()
							//log.Printf("r.Closed")
						}()
						//body.StreamChan = make(chan *ChatStreamChunk, 128)
						const dataField = "data: "
						scanner := bufio.NewScanner(r)
						for scanner.Scan() {
							line := scanner.Bytes()
							if bytes.HasPrefix(line, []byte(dataField)) {
								line = bytes.TrimPrefix(line, []byte(dataField))
								if bytes.Equal(line, []byte("[DONE]")) {
									//log.Printf("close: DONE")
									close(b.StreamChan)
									break
								}
								var chunk ChatStreamChunk
								if err = json.Unmarshal(line, &chunk); err != nil {
									//
									log.Printf("unmarshal error: %v", err)
								} else {
									b.StreamChan <- &chunk
								}
								//log.Printf("data: [%s]", line)
							} else {
								//log.Printf("line: [%s]", line)
							}
						}
					}(b, res.Body)
				}
				return nil
			}
		}
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}
