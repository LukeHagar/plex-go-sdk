// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type DeleteLibraryLibraryErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *DeleteLibraryLibraryErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *DeleteLibraryLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteLibraryLibraryErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// DeleteLibraryLibraryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type DeleteLibraryLibraryResponseBody struct {
	Errors []DeleteLibraryLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &DeleteLibraryLibraryResponseBody{}

func (e *DeleteLibraryLibraryResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DeleteLibraryErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *DeleteLibraryErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *DeleteLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteLibraryErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// DeleteLibraryResponseBody - Bad Request - A parameter was not specified, or was specified incorrectly.
type DeleteLibraryResponseBody struct {
	Errors []DeleteLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &DeleteLibraryResponseBody{}

func (e *DeleteLibraryResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
