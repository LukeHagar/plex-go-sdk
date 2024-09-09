// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetAllLibrariesLibraryErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetAllLibrariesLibraryErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetAllLibrariesLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetAllLibrariesLibraryErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetAllLibrariesLibraryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetAllLibrariesLibraryResponseBody struct {
	Errors []GetAllLibrariesLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetAllLibrariesLibraryResponseBody{}

func (e *GetAllLibrariesLibraryResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetAllLibrariesErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetAllLibrariesErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetAllLibrariesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetAllLibrariesErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetAllLibrariesResponseBody - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetAllLibrariesResponseBody struct {
	Errors []GetAllLibrariesErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetAllLibrariesResponseBody{}

func (e *GetAllLibrariesResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
