// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetTokenDetailsAuthenticationErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetTokenDetailsAuthenticationErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTokenDetailsAuthenticationErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTokenDetailsAuthenticationErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTokenDetailsUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetTokenDetailsUnauthorized struct {
	Errors []GetTokenDetailsAuthenticationErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetTokenDetailsUnauthorized{}

func (e *GetTokenDetailsUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetTokenDetailsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetTokenDetailsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTokenDetailsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTokenDetailsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTokenDetailsBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetTokenDetailsBadRequest struct {
	Errors []GetTokenDetailsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetTokenDetailsBadRequest{}

func (e *GetTokenDetailsBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
