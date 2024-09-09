// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"net/http"
)

// GetServerIdentityResponseBody - Request Timeout
type GetServerIdentityResponseBody struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetServerIdentityResponseBody{}

func (e *GetServerIdentityResponseBody) Error() string {
	if e.Message == nil {
		return "unknown error"
	}

	return *e.Message
}
