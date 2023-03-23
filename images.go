package openai

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const (
	Size256  = "256x256"
	Size512  = "512x512"
	Size1024 = "1024x1024"
)

const (
	ImageResponseFormat  = "url"
	ImageResponseB64Json = "b64_json"
)

type ImageRequestBody struct {
	// [Required]
	// A text description of the desired image(s).
	// The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// [Optional Defaults to 1]
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `json:"n,omitempty"`
	// [Optional Defaults to 1024x1024] The size of the generated images.
	//Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size string `json:"size,omitempty"`
	// [Optional Defaults to url]
	// The format in which the generated images are returned.
	// Must be one of `url` or `b64_json`.
	ResponseFormat string `json:"response_format,omitempty"`
	// [Optional]
	// A unique identifier representing your end-user,
	// which can help OpenAI to monitor and detect abuse.
	User string `json:"user,omitempty"`
}

type ImageEditRequestBody struct {
	// [Required]
	// The image to edit. Must be a valid PNG file, less than 4MB, and square.
	// If mask is not provided, image must have transparency, which will be used as the mask.
	Image string `json:"image" multipart:"image,file"`
	// [Optional]
	// An additional image whose fully transparent areas (e.g. where alpha is zero) indicate
	// where `image` should be edited. Must be a valid PNG file, less than 4MB, and have the
	// same dimensions as `image`.
	Mask string `json:"mask,omitempty" multipart:"mask,file"`
	// [Required]
	// A text description of the desired image(s).
	// The maximum length is 1000 characters.
	Prompt string `json:"prompt" multipart:"prompt"`
	// [Optional Defaults to 1]
	// The number of images to generate.
	// Must be between 1 and 10.
	N int `json:"n,omitempty" multipart:"n"`
	// [Optional Defaults to 1024x1024]
	// The size of the generated images.
	// Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size string `json:"size,omitempty" multipart:"size"`
	// [Optional Defaults to url]
	// The format in which the generated images are returned.
	// Must be one of `url` or `b64_json`.
	ResponseFormat string `json:"response_format,omitempty" multipart:"response-format"`
	// [Optional]
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	User string `json:"user,omitempty" multipart:"user"`
}

func (b ImageEditRequestBody) WriteForm(w *multipart.Writer) (err error) {
	var imageWriter io.Writer
	if imageWriter, err = w.CreateFormFile("image", b.Image); err != nil {
		return
	}

	var imageFile *os.File
	if imageFile, err = os.Open(b.Image); err != nil {
		return
	}

	defer func() {
		_ = imageFile.Close()
	}()

	if _, err = io.Copy(imageWriter, imageFile); err != nil {
		return
	}

	if b.Mask != "" {
		var maskWriter io.Writer
		if maskWriter, err = w.CreateFormFile("mask", b.Mask); err != nil {
			return
		}
		var maskFile *os.File
		if maskFile, err = os.Open(b.Mask); err != nil {
			return
		}
		defer func() { _ = maskFile.Close() }()
		if _, err = io.Copy(maskWriter, maskFile); err != nil {
			return
		}
	}

	if b.Prompt != "" {
		if err = w.WriteField("prompt", b.Prompt); err != nil {
			return
		}
	}

	if err = w.WriteField("n", strconv.Itoa(b.N)); err != nil {
		return
	}

	if b.Size != "" {
		if err = w.WriteField("size", b.Size); err != nil {
			return
		}
	}

	return
}

type ImageVariationRequestBody struct {
	// [Required]
	// The image to use as the basic for the variation(s).
	// Must be a valid PNG file, less than 4MB, and square.
	Image          string `json:"image"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

func (b ImageVariationRequestBody) WriteForm(w *multipart.Writer) (err error) {
	imageWriter, err := w.CreateFormFile("image", b.Image)
	if err != nil {
		return
	}

	var f *os.File
	if f, err = os.Open(b.Image); err != nil {
		return
	}

	defer func() {
		_ = f.Close()
	}()

	if _, err = io.Copy(imageWriter, f); err != nil {
		return
	}

	if err = w.WriteField("n", strconv.Itoa(b.N)); err != nil {
		return
	}

	if b.Size != "" {
		if err = w.WriteField("size", b.Size); err != nil {
			return
		}
	}
	if err = w.WriteField("n", strconv.Itoa(b.N)); err != nil {
		return
	}

	if b.Size != "" {
		if err = w.WriteField("size", b.Size); err != nil {
			return
		}
	}

	if b.ResponseFormat != "" {
		if err = w.WriteField("response_format", b.ResponseFormat); err != nil {
			return
		}
	}

	if b.User != "" {
		if err = w.WriteField("user", b.User); err != nil {
			return
		}
	}

	return
}

type ImageData struct {
	URL string `json:"url"`
}

type ImageResponseBody struct {
	Created int         `json:"created"`
	Data    []ImageData `json:"data"`
}

// CreateImage Creates an image given a prompt.
// POST https://api.openai.com/v1/images/generations
func (c *Client) CreateImage(
	ctx context.Context,
	reqBody ImageRequestBody) (resBody ImageResponseBody, err error) {
	const apiURL = "https://api.openai.com/v1/images/generations"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}
	err = c.getRequest(req, &resBody)
	return
}

// CreateImageEdit Create an edited or extended image given an original image and a prompt.
// POST https://api.openai.com/v1/images/edits
func (c *Client) CreateImageEdit(
	ctx context.Context,
	reqBody ImageEditRequestBody) (resBody ImageResponseBody, err error) {
	const apiURL = "https://api.openai.com/v1/images/edits"

	switch reqBody.Size {
	case Size256, Size512, Size1024, "":
	default:
		err = errors.New("invalid `size`")
		return
	}

	switch reqBody.ResponseFormat {
	case ImageResponseFormat, ImageResponseB64Json, "":
	default:
		err = errors.New("invalid `response_format`")
		return
	}

	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}
	err = c.getRequest(req, &resBody)
	return
}

// CreateImageVariation Create a variation of a given image.
// POST https://api.openai.com/v1/images/variations
func (c *Client) CreateImageVariation(
	ctx context.Context,
	reqBody ImageVariationRequestBody) (resBody ImageResponseBody, err error) {
	const apiURL = "https://api.openai.com/v1/images/variations"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}
	err = c.getRequest(req, &resBody)
	return
}
