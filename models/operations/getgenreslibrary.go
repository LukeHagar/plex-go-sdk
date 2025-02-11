// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetGenresLibraryQueryParamType - The type of media to retrieve or filter by.
// 1 = movie
// 2 = show
// 3 = season
// 4 = episode
// E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries
type GetGenresLibraryQueryParamType int64

const (
	GetGenresLibraryQueryParamTypeMovie   GetGenresLibraryQueryParamType = 1
	GetGenresLibraryQueryParamTypeTvShow  GetGenresLibraryQueryParamType = 2
	GetGenresLibraryQueryParamTypeSeason  GetGenresLibraryQueryParamType = 3
	GetGenresLibraryQueryParamTypeEpisode GetGenresLibraryQueryParamType = 4
	GetGenresLibraryQueryParamTypeAudio   GetGenresLibraryQueryParamType = 8
	GetGenresLibraryQueryParamTypeAlbum   GetGenresLibraryQueryParamType = 9
	GetGenresLibraryQueryParamTypeTrack   GetGenresLibraryQueryParamType = 10
)

func (e GetGenresLibraryQueryParamType) ToPointer() *GetGenresLibraryQueryParamType {
	return &e
}
func (e *GetGenresLibraryQueryParamType) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 10:
		*e = GetGenresLibraryQueryParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetGenresLibraryQueryParamType: %v", v)
	}
}

type GetGenresLibraryRequest struct {
	// The unique key of the Plex library.
	// Note: This is unique in the context of the Plex server.
	//
	SectionKey int `pathParam:"style=simple,explode=false,name=sectionKey"`
	// The type of media to retrieve or filter by.
	// 1 = movie
	// 2 = show
	// 3 = season
	// 4 = episode
	// E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries
	//
	Type GetGenresLibraryQueryParamType `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetGenresLibraryRequest) GetSectionKey() int {
	if o == nil {
		return 0
	}
	return o.SectionKey
}

func (o *GetGenresLibraryRequest) GetType() GetGenresLibraryQueryParamType {
	if o == nil {
		return GetGenresLibraryQueryParamType(0)
	}
	return o.Type
}

type GetGenresLibraryDirectory struct {
	FastKey string `json:"fastKey"`
	Key     string `json:"key"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

func (o *GetGenresLibraryDirectory) GetFastKey() string {
	if o == nil {
		return ""
	}
	return o.FastKey
}

func (o *GetGenresLibraryDirectory) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *GetGenresLibraryDirectory) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *GetGenresLibraryDirectory) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type GetGenresLibraryMediaContainer struct {
	// Number of media items returned in this response.
	Size int `json:"size"`
	// Indicates whether syncing is allowed.
	AllowSync bool `json:"allowSync"`
	// URL for the background artwork of the media container.
	Art string `json:"art"`
	// The content type or mode.
	Content string `json:"content"`
	// An plugin identifier for the media container.
	Identifier string `json:"identifier"`
	// The prefix used for media tag resource paths.
	MediaTagPrefix string `json:"mediaTagPrefix"`
	// The version number for media tags.
	MediaTagVersion int64 `json:"mediaTagVersion"`
	// Specifies whether caching is disabled.
	Nocache bool `json:"nocache"`
	// URL for the thumbnail image of the media container.
	Thumb string `json:"thumb"`
	// The primary title of the media container.
	Title1 string `json:"title1"`
	// The secondary title of the media container.
	Title2 string `json:"title2"`
	// Identifier for the view group layout.
	ViewGroup string                      `json:"viewGroup"`
	Directory []GetGenresLibraryDirectory `json:"Directory,omitempty"`
}

func (o *GetGenresLibraryMediaContainer) GetSize() int {
	if o == nil {
		return 0
	}
	return o.Size
}

func (o *GetGenresLibraryMediaContainer) GetAllowSync() bool {
	if o == nil {
		return false
	}
	return o.AllowSync
}

func (o *GetGenresLibraryMediaContainer) GetArt() string {
	if o == nil {
		return ""
	}
	return o.Art
}

func (o *GetGenresLibraryMediaContainer) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *GetGenresLibraryMediaContainer) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *GetGenresLibraryMediaContainer) GetMediaTagPrefix() string {
	if o == nil {
		return ""
	}
	return o.MediaTagPrefix
}

func (o *GetGenresLibraryMediaContainer) GetMediaTagVersion() int64 {
	if o == nil {
		return 0
	}
	return o.MediaTagVersion
}

func (o *GetGenresLibraryMediaContainer) GetNocache() bool {
	if o == nil {
		return false
	}
	return o.Nocache
}

func (o *GetGenresLibraryMediaContainer) GetThumb() string {
	if o == nil {
		return ""
	}
	return o.Thumb
}

func (o *GetGenresLibraryMediaContainer) GetTitle1() string {
	if o == nil {
		return ""
	}
	return o.Title1
}

func (o *GetGenresLibraryMediaContainer) GetTitle2() string {
	if o == nil {
		return ""
	}
	return o.Title2
}

func (o *GetGenresLibraryMediaContainer) GetViewGroup() string {
	if o == nil {
		return ""
	}
	return o.ViewGroup
}

func (o *GetGenresLibraryMediaContainer) GetDirectory() []GetGenresLibraryDirectory {
	if o == nil {
		return nil
	}
	return o.Directory
}

// GetGenresLibraryResponseBody - Successful response containing media container data.
type GetGenresLibraryResponseBody struct {
	MediaContainer *GetGenresLibraryMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetGenresLibraryResponseBody) GetMediaContainer() *GetGenresLibraryMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetGenresLibraryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response containing media container data.
	Object *GetGenresLibraryResponseBody
}

func (o *GetGenresLibraryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGenresLibraryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGenresLibraryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGenresLibraryResponse) GetObject() *GetGenresLibraryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
