// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetHomeDataPlexErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetHomeDataPlexErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetHomeDataPlexErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetHomeDataPlexErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetHomeDataPlexResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetHomeDataPlexResponseBody struct {
	Errors []GetHomeDataPlexErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetHomeDataPlexResponseBody{}

func (e *GetHomeDataPlexResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetHomeDataErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetHomeDataErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetHomeDataErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetHomeDataErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetHomeDataResponseBody - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetHomeDataResponseBody struct {
	Errors []GetHomeDataErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetHomeDataResponseBody{}

func (e *GetHomeDataResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
