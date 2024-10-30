// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/models/shared"
	"net/http"
)

type RetrieveFileRequest struct {
	// The ID of the file to use for this request
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

func (o *RetrieveFileRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

type RetrieveFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	OpenAIFile *shared.OpenAIFile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveFileResponse) GetOpenAIFile() *shared.OpenAIFile {
	if o == nil {
		return nil
	}
	return o.OpenAIFile
}

func (o *RetrieveFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
