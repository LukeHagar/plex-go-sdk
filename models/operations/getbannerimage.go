// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type GetBannerImageRequest struct {
	// the id of the library item to return the children of.
	RatingKey int64 `pathParam:"style=simple,explode=false,name=ratingKey"`
	Width     int64 `queryParam:"style=form,explode=true,name=width"`
	Height    int64 `queryParam:"style=form,explode=true,name=height"`
	MinSize   int64 `queryParam:"style=form,explode=true,name=minSize"`
	Upscale   int64 `queryParam:"style=form,explode=true,name=upscale"`
	// Plex Authentication Token
	XPlexToken string `queryParam:"style=form,explode=true,name=X-Plex-Token"`
}

func (o *GetBannerImageRequest) GetRatingKey() int64 {
	if o == nil {
		return 0
	}
	return o.RatingKey
}

func (o *GetBannerImageRequest) GetWidth() int64 {
	if o == nil {
		return 0
	}
	return o.Width
}

func (o *GetBannerImageRequest) GetHeight() int64 {
	if o == nil {
		return 0
	}
	return o.Height
}

func (o *GetBannerImageRequest) GetMinSize() int64 {
	if o == nil {
		return 0
	}
	return o.MinSize
}

func (o *GetBannerImageRequest) GetUpscale() int64 {
	if o == nil {
		return 0
	}
	return o.Upscale
}

func (o *GetBannerImageRequest) GetXPlexToken() string {
	if o == nil {
		return ""
	}
	return o.XPlexToken
}

type GetBannerImageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response returning an image
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	ResponseStream io.ReadCloser
	Headers        map[string][]string
}

func (o *GetBannerImageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBannerImageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBannerImageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetBannerImageResponse) GetResponseStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.ResponseStream
}

func (o *GetBannerImageResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
