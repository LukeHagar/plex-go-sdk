// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetStatisticsStatisticsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetStatisticsStatisticsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetStatisticsStatisticsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetStatisticsStatisticsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetStatisticsUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetStatisticsUnauthorized struct {
	Errors []GetStatisticsStatisticsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetStatisticsUnauthorized{}

func (e *GetStatisticsUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetStatisticsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetStatisticsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetStatisticsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetStatisticsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetStatisticsBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetStatisticsBadRequest struct {
	Errors []GetStatisticsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetStatisticsBadRequest{}

func (e *GetStatisticsBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
