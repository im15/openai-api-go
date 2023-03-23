package openai

import (
	"context"
	"net/http"
)

const (
	GPT4                 = "gpt-4"
	GPT40314             = "gpt-4-0314"
	GPT432k              = "gpt-4-32k"
	GPT432k0314          = "gpt-4-32k-0314"
	GPT35Turbo           = "gpt-3.5-turbo"
	GPT35Turbo0310       = "gpt-3.5-turbo-0310"
	TextDavinci003       = "text-davinci-003"
	TextDavinci002       = "text-davinci-002"
	TextCurie001         = "text-curie-001"
	TextBabBage001       = "text-babbage-001"
	TextAda001           = "text-ada-001"
	Davinci              = "davinci"
	Curie                = "curie"
	Babbage              = "babbage"
	Ada                  = "ada"
	TextDavinciEdit001   = "text-davinci-edit-001"
	CodeDavinciEdit001   = "code-davinci-edit-001"
	Whisper1             = "whisper-1"
	TextEmbeddingAda002  = "text-embedding-ada-002"
	TextSearchAdaDoc001  = "text-search-ada-doc-001"
	TextModerationStable = "text-moderation-stable"
	TextModerationLatest = "text-moderation-latest"
)

type ModelPermission struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Created            int         `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}

type ModelObject struct {
	ID         string            `json:"id"`
	Object     string            `json:"object"`
	Created    int               `json:"created"`
	OwnerBy    string            `json:"owner_by"`
	Permission []ModelPermission `json:"permission"`
	Root       string            `json:"root"`
	Parent     interface{}       `json:"parent"`
}

type ModelsResponseBody struct {
	Object string        `json:"object"`
	Data   []ModelObject `json:"data"`
}

// ListModels Lists the currently available models, and provides basic information about each
// one such as the owner and availability.
func (c *Client) ListModels(ctx context.Context) (*ModelsResponseBody, error) {
	const apiUrlV1 = "https://api.openai.com/v1/models"
	req, err := c.newRequest(ctx, GET, apiUrlV1, nil)
	if err != nil {
		return nil, err
	}
	var body ModelsResponseBody
	if err := c.getRequest(req, &body); err != nil {
		return nil, err
	}
	return &body, nil
}

// RetrieveModel Retrieves a model instance, providing basic information about the model
// such as the owner and permissioning.
// `model`: The ID of the model to use for this request
func (c *Client) RetrieveModel(ctx context.Context, model string) (modelObject *ModelObject, err error) {
	const apiURLv1 = "https://api.openai.com/v1/models/"
	var req *http.Request
	if req, err = c.newRequest(ctx, GET, apiURLv1+model, nil); err != nil {
		return nil, err
	}

	var data ModelObject
	if err = c.getRequest(req, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
