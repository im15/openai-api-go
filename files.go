package openai

import (
	"context"
	"fmt"
	"net/http"
)

type FileObject struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type UploadFileRequestBody struct {
	File    string `json:"file"`
	Purpose string `json:"purpose"`
}

type ListFilesResponseBody struct {
	Object string       `json:"object"`
	Data   []FileObject `json:"data"`
}

type DeleteFileResponseBody struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type RetrieveFileContentResponseBody struct {
	Data []byte
}

// ListFiles Return a list of files that belong to the user's organization.
// GET https://api.openai.com/v1/files
func (c *Client) ListFiles(ctx context.Context) (resBody ListFilesResponseBody, err error) {
	const apiURL = apiURLPrefix + "/v1/files"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodGet, apiURL, nil); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}

// UploadFile Upload a file that contains document(s) to be used across various
// endpoints/features. Currently, the size of all the files uploaded by one organization
// can be up to 1GB. Please contact us if you need to increase the storage limit.
// POST https://api.openai.com/v1/files
func (c *Client) UploadFile(
	ctx context.Context,
	reqBody UploadFileRequestBody) (resBody FileObject, err error) {
	const apiURL = apiURLPrefix + "/v1/files"
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodPost, apiURL, reqBody); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}

// DeleteFile Delete a file
// DELETE https://api.openai.com/v1/files/{file_id}
func (c *Client) DeleteFile(
	ctx context.Context,
	fileID string) (resBody DeleteFileResponseBody, err error) {
	var apiURL = fmt.Sprintf("%s/v1/files/%s", apiURLPrefix, fileID)

	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodDelete, apiURL, nil); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}

// RetrieveFile Return information about a specific file.
// GET https://api.openai.com/v1/files/{file_id}
func (c *Client) RetrieveFile(
	ctx context.Context,
	fileID string) (resBody FileObject, err error) {
	var apiURL = fmt.Sprintf("%s/v1/files/%s", apiURLPrefix, fileID)
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodGet, apiURL, nil); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}

// RetrieveFileContent Returns the contents of the specified file.
// GET https://api.openai.com/v1/files/{file_id}/content
func (c *Client) RetrieveFileContent(
	ctx context.Context,
	fileID string) (resBody RetrieveFileContentResponseBody, err error) {
	var apiURL = fmt.Sprintf("%s/v1/files/%s/content", apiURLPrefix, fileID)
	var req *http.Request
	if req, err = c.newRequest(ctx, http.MethodGet, apiURL, nil); err != nil {
		return
	}

	err = c.getRequest(req, &resBody)

	return
}
