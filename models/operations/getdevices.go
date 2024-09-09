// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type Device struct {
	ID               *float64 `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Platform         *string  `json:"platform,omitempty"`
	ClientIdentifier *string  `json:"clientIdentifier,omitempty"`
	CreatedAt        *float64 `json:"createdAt,omitempty"`
}

func (o *Device) GetID() *float64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Device) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Device) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *Device) GetClientIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ClientIdentifier
}

func (o *Device) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

type GetDevicesMediaContainer struct {
	Size       *float64 `json:"size,omitempty"`
	Identifier *string  `json:"identifier,omitempty"`
	Device     []Device `json:"Device,omitempty"`
}

func (o *GetDevicesMediaContainer) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetDevicesMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetDevicesMediaContainer) GetDevice() []Device {
	if o == nil {
		return nil
	}
	return o.Device
}

// GetDevicesResponseBody - Devices
type GetDevicesResponseBody struct {
	MediaContainer *GetDevicesMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetDevicesResponseBody) GetMediaContainer() *GetDevicesMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetDevicesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Devices
	Object *GetDevicesResponseBody
}

func (o *GetDevicesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDevicesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDevicesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDevicesResponse) GetObject() *GetDevicesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
