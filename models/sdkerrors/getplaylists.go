// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetPlaylistsPlaylistsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetPlaylistsPlaylistsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetPlaylistsPlaylistsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPlaylistsPlaylistsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetPlaylistsUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetPlaylistsUnauthorized struct {
	Errors []GetPlaylistsPlaylistsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetPlaylistsUnauthorized{}

func (e *GetPlaylistsUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetPlaylistsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetPlaylistsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetPlaylistsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPlaylistsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetPlaylistsBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetPlaylistsBadRequest struct {
	Errors []GetPlaylistsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetPlaylistsBadRequest{}

func (e *GetPlaylistsBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
