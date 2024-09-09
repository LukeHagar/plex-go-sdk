// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetServerIdentityMediaContainer struct {
	Size              *float64 `json:"size,omitempty"`
	Claimed           *bool    `json:"claimed,omitempty"`
	MachineIdentifier *string  `json:"machineIdentifier,omitempty"`
	Version           *string  `json:"version,omitempty"`
}

func (o *GetServerIdentityMediaContainer) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetServerIdentityMediaContainer) GetClaimed() *bool {
	if o == nil {
		return nil
	}
	return o.Claimed
}

func (o *GetServerIdentityMediaContainer) GetMachineIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.MachineIdentifier
}

func (o *GetServerIdentityMediaContainer) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

// GetServerIdentityResponseBody - The Server Identity information
type GetServerIdentityResponseBody struct {
	MediaContainer *GetServerIdentityMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetServerIdentityResponseBody) GetMediaContainer() *GetServerIdentityMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetServerIdentityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The Server Identity information
	Object *GetServerIdentityResponseBody
}

func (o *GetServerIdentityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerIdentityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerIdentityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServerIdentityResponse) GetObject() *GetServerIdentityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
