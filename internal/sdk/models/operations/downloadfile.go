// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DownloadFileRequest struct {
	// The ID of the file to use for this request
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

func (o *DownloadFileRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

type DownloadFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	String *string
}

func (o *DownloadFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DownloadFileResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}
