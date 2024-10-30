// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/models/shared"
	"net/http"
)

type CreateImageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ImagesResponse *shared.ImagesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateImageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateImageResponse) GetImagesResponse() *shared.ImagesResponse {
	if o == nil {
		return nil
	}
	return o.ImagesResponse
}

func (o *CreateImageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateImageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
