// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type ClearPlaylistContentsPlaylistsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *ClearPlaylistContentsPlaylistsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ClearPlaylistContentsPlaylistsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ClearPlaylistContentsPlaylistsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// ClearPlaylistContentsUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type ClearPlaylistContentsUnauthorized struct {
	Errors []ClearPlaylistContentsPlaylistsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &ClearPlaylistContentsUnauthorized{}

func (e *ClearPlaylistContentsUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type ClearPlaylistContentsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *ClearPlaylistContentsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ClearPlaylistContentsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ClearPlaylistContentsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// ClearPlaylistContentsBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type ClearPlaylistContentsBadRequest struct {
	Errors []ClearPlaylistContentsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &ClearPlaylistContentsBadRequest{}

func (e *ClearPlaylistContentsBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
