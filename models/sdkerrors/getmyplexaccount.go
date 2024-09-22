// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetMyPlexAccountServerErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetMyPlexAccountServerErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMyPlexAccountServerErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMyPlexAccountServerErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMyPlexAccountUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetMyPlexAccountUnauthorized struct {
	Errors []GetMyPlexAccountServerErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMyPlexAccountUnauthorized{}

func (e *GetMyPlexAccountUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetMyPlexAccountErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetMyPlexAccountErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMyPlexAccountErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMyPlexAccountErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMyPlexAccountBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetMyPlexAccountBadRequest struct {
	Errors []GetMyPlexAccountErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMyPlexAccountBadRequest{}

func (e *GetMyPlexAccountBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
