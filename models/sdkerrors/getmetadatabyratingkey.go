// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetMetaDataByRatingKeyLibraryErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetMetaDataByRatingKeyLibraryErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMetaDataByRatingKeyLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMetaDataByRatingKeyLibraryErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMetaDataByRatingKeyUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetMetaDataByRatingKeyUnauthorized struct {
	Errors []GetMetaDataByRatingKeyLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMetaDataByRatingKeyUnauthorized{}

func (e *GetMetaDataByRatingKeyUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetMetaDataByRatingKeyErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetMetaDataByRatingKeyErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMetaDataByRatingKeyErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMetaDataByRatingKeyErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMetaDataByRatingKeyBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetMetaDataByRatingKeyBadRequest struct {
	Errors []GetMetaDataByRatingKeyErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMetaDataByRatingKeyBadRequest{}

func (e *GetMetaDataByRatingKeyBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
