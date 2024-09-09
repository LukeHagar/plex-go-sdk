// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetGeoDataPlexErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetGeoDataPlexErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetGeoDataPlexErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetGeoDataPlexErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetGeoDataPlexResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetGeoDataPlexResponseBody struct {
	Errors []GetGeoDataPlexErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetGeoDataPlexResponseBody{}

func (e *GetGeoDataPlexResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetGeoDataErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetGeoDataErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetGeoDataErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetGeoDataErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetGeoDataResponseBody - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetGeoDataResponseBody struct {
	Errors []GetGeoDataErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetGeoDataResponseBody{}

func (e *GetGeoDataResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
