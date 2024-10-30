// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateTranscriptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateTranscriptionResponse *shared.CreateTranscriptionResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateTranscriptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTranscriptionResponse) GetCreateTranscriptionResponse() *shared.CreateTranscriptionResponse {
	if o == nil {
		return nil
	}
	return o.CreateTranscriptionResponse
}

func (o *CreateTranscriptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTranscriptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
