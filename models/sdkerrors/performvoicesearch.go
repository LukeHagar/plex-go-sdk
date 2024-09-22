// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type PerformVoiceSearchSearchErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *PerformVoiceSearchSearchErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *PerformVoiceSearchSearchErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PerformVoiceSearchSearchErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// PerformVoiceSearchUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type PerformVoiceSearchUnauthorized struct {
	Errors []PerformVoiceSearchSearchErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PerformVoiceSearchUnauthorized{}

func (e *PerformVoiceSearchUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type PerformVoiceSearchErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *PerformVoiceSearchErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *PerformVoiceSearchErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PerformVoiceSearchErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// PerformVoiceSearchBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type PerformVoiceSearchBadRequest struct {
	Errors []PerformVoiceSearchErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PerformVoiceSearchBadRequest{}

func (e *PerformVoiceSearchBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
