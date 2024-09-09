// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type DeletePlaylistPlaylistsErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *DeletePlaylistPlaylistsErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *DeletePlaylistPlaylistsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeletePlaylistPlaylistsErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// DeletePlaylistPlaylistsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type DeletePlaylistPlaylistsResponseBody struct {
	Errors []DeletePlaylistPlaylistsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &DeletePlaylistPlaylistsResponseBody{}

func (e *DeletePlaylistPlaylistsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type DeletePlaylistErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *DeletePlaylistErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *DeletePlaylistErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeletePlaylistErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// DeletePlaylistResponseBody - Bad Request - A parameter was not specified, or was specified incorrectly.
type DeletePlaylistResponseBody struct {
	Errors []DeletePlaylistErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &DeletePlaylistResponseBody{}

func (e *DeletePlaylistResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
