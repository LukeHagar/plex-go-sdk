// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetMetadataChildrenRequest struct {
	// the id of the library item to return the children of.
	RatingKey float64 `pathParam:"style=simple,explode=false,name=ratingKey"`
	// Adds additional elements to the response. Supported types are (Stream)
	//
	IncludeElements *string `queryParam:"style=form,explode=true,name=includeElements"`
}

func (o *GetMetadataChildrenRequest) GetRatingKey() float64 {
	if o == nil {
		return 0.0
	}
	return o.RatingKey
}

func (o *GetMetadataChildrenRequest) GetIncludeElements() *string {
	if o == nil {
		return nil
	}
	return o.IncludeElements
}

type GetMetadataChildrenDirectory struct {
	LeafCount       *int    `json:"leafCount,omitempty"`
	Thumb           *string `json:"thumb,omitempty"`
	ViewedLeafCount *int    `json:"viewedLeafCount,omitempty"`
	Key             *string `json:"key,omitempty"`
	Title           *string `json:"title,omitempty"`
}

func (o *GetMetadataChildrenDirectory) GetLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *GetMetadataChildrenDirectory) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetMetadataChildrenDirectory) GetViewedLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.ViewedLeafCount
}

func (o *GetMetadataChildrenDirectory) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMetadataChildrenDirectory) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type GetMetadataChildrenMetadata struct {
	RatingKey       *string `json:"ratingKey,omitempty"`
	Key             *string `json:"key,omitempty"`
	ParentRatingKey *string `json:"parentRatingKey,omitempty"`
	GUID            *string `json:"guid,omitempty"`
	ParentGUID      *string `json:"parentGuid,omitempty"`
	ParentStudio    *string `json:"parentStudio,omitempty"`
	Type            *string `json:"type,omitempty"`
	Title           *string `json:"title,omitempty"`
	ParentKey       *string `json:"parentKey,omitempty"`
	ParentTitle     *string `json:"parentTitle,omitempty"`
	Summary         *string `json:"summary,omitempty"`
	Index           *int    `json:"index,omitempty"`
	ParentIndex     *int    `json:"parentIndex,omitempty"`
	ViewCount       *int    `json:"viewCount,omitempty"`
	LastViewedAt    *int    `json:"lastViewedAt,omitempty"`
	ParentYear      *int    `json:"parentYear,omitempty"`
	Thumb           *string `json:"thumb,omitempty"`
	Art             *string `json:"art,omitempty"`
	ParentThumb     *string `json:"parentThumb,omitempty"`
	ParentTheme     *string `json:"parentTheme,omitempty"`
	LeafCount       *int    `json:"leafCount,omitempty"`
	ViewedLeafCount *int    `json:"viewedLeafCount,omitempty"`
	AddedAt         *int    `json:"addedAt,omitempty"`
	UpdatedAt       *int    `json:"updatedAt,omitempty"`
	UserRating      *int    `json:"userRating,omitempty"`
	SkipCount       *int    `json:"skipCount,omitempty"`
	LastRatedAt     *int    `json:"lastRatedAt,omitempty"`
}

func (o *GetMetadataChildrenMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetMetadataChildrenMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMetadataChildrenMetadata) GetParentRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.ParentRatingKey
}

func (o *GetMetadataChildrenMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetMetadataChildrenMetadata) GetParentGUID() *string {
	if o == nil {
		return nil
	}
	return o.ParentGUID
}

func (o *GetMetadataChildrenMetadata) GetParentStudio() *string {
	if o == nil {
		return nil
	}
	return o.ParentStudio
}

func (o *GetMetadataChildrenMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetMetadataChildrenMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetMetadataChildrenMetadata) GetParentKey() *string {
	if o == nil {
		return nil
	}
	return o.ParentKey
}

func (o *GetMetadataChildrenMetadata) GetParentTitle() *string {
	if o == nil {
		return nil
	}
	return o.ParentTitle
}

func (o *GetMetadataChildrenMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetMetadataChildrenMetadata) GetIndex() *int {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetMetadataChildrenMetadata) GetParentIndex() *int {
	if o == nil {
		return nil
	}
	return o.ParentIndex
}

func (o *GetMetadataChildrenMetadata) GetViewCount() *int {
	if o == nil {
		return nil
	}
	return o.ViewCount
}

func (o *GetMetadataChildrenMetadata) GetLastViewedAt() *int {
	if o == nil {
		return nil
	}
	return o.LastViewedAt
}

func (o *GetMetadataChildrenMetadata) GetParentYear() *int {
	if o == nil {
		return nil
	}
	return o.ParentYear
}

func (o *GetMetadataChildrenMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetMetadataChildrenMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetMetadataChildrenMetadata) GetParentThumb() *string {
	if o == nil {
		return nil
	}
	return o.ParentThumb
}

func (o *GetMetadataChildrenMetadata) GetParentTheme() *string {
	if o == nil {
		return nil
	}
	return o.ParentTheme
}

func (o *GetMetadataChildrenMetadata) GetLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *GetMetadataChildrenMetadata) GetViewedLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.ViewedLeafCount
}

func (o *GetMetadataChildrenMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetMetadataChildrenMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetMetadataChildrenMetadata) GetUserRating() *int {
	if o == nil {
		return nil
	}
	return o.UserRating
}

func (o *GetMetadataChildrenMetadata) GetSkipCount() *int {
	if o == nil {
		return nil
	}
	return o.SkipCount
}

func (o *GetMetadataChildrenMetadata) GetLastRatedAt() *int {
	if o == nil {
		return nil
	}
	return o.LastRatedAt
}

type GetMetadataChildrenMediaContainer struct {
	Size                *int                           `json:"size,omitempty"`
	AllowSync           *bool                          `json:"allowSync,omitempty"`
	Art                 *string                        `json:"art,omitempty"`
	Identifier          *string                        `json:"identifier,omitempty"`
	Key                 *string                        `json:"key,omitempty"`
	LibrarySectionID    *int                           `json:"librarySectionID,omitempty"`
	LibrarySectionTitle *string                        `json:"librarySectionTitle,omitempty"`
	LibrarySectionUUID  *string                        `json:"librarySectionUUID,omitempty"`
	MediaTagPrefix      *string                        `json:"mediaTagPrefix,omitempty"`
	MediaTagVersion     *int                           `json:"mediaTagVersion,omitempty"`
	Nocache             *bool                          `json:"nocache,omitempty"`
	ParentIndex         *int                           `json:"parentIndex,omitempty"`
	ParentTitle         *string                        `json:"parentTitle,omitempty"`
	ParentYear          *int                           `json:"parentYear,omitempty"`
	Summary             *string                        `json:"summary,omitempty"`
	Theme               *string                        `json:"theme,omitempty"`
	Thumb               *string                        `json:"thumb,omitempty"`
	Title1              *string                        `json:"title1,omitempty"`
	Title2              *string                        `json:"title2,omitempty"`
	ViewGroup           *string                        `json:"viewGroup,omitempty"`
	ViewMode            *int                           `json:"viewMode,omitempty"`
	Directory           []GetMetadataChildrenDirectory `json:"Directory,omitempty"`
	Metadata            []GetMetadataChildrenMetadata  `json:"Metadata,omitempty"`
}

func (o *GetMetadataChildrenMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetMetadataChildrenMediaContainer) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetMetadataChildrenMediaContainer) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetMetadataChildrenMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetMetadataChildrenMediaContainer) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMetadataChildrenMediaContainer) GetLibrarySectionID() *int {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetMetadataChildrenMediaContainer) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetMetadataChildrenMediaContainer) GetLibrarySectionUUID() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionUUID
}

func (o *GetMetadataChildrenMediaContainer) GetMediaTagPrefix() *string {
	if o == nil {
		return nil
	}
	return o.MediaTagPrefix
}

func (o *GetMetadataChildrenMediaContainer) GetMediaTagVersion() *int {
	if o == nil {
		return nil
	}
	return o.MediaTagVersion
}

func (o *GetMetadataChildrenMediaContainer) GetNocache() *bool {
	if o == nil {
		return nil
	}
	return o.Nocache
}

func (o *GetMetadataChildrenMediaContainer) GetParentIndex() *int {
	if o == nil {
		return nil
	}
	return o.ParentIndex
}

func (o *GetMetadataChildrenMediaContainer) GetParentTitle() *string {
	if o == nil {
		return nil
	}
	return o.ParentTitle
}

func (o *GetMetadataChildrenMediaContainer) GetParentYear() *int {
	if o == nil {
		return nil
	}
	return o.ParentYear
}

func (o *GetMetadataChildrenMediaContainer) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetMetadataChildrenMediaContainer) GetTheme() *string {
	if o == nil {
		return nil
	}
	return o.Theme
}

func (o *GetMetadataChildrenMediaContainer) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetMetadataChildrenMediaContainer) GetTitle1() *string {
	if o == nil {
		return nil
	}
	return o.Title1
}

func (o *GetMetadataChildrenMediaContainer) GetTitle2() *string {
	if o == nil {
		return nil
	}
	return o.Title2
}

func (o *GetMetadataChildrenMediaContainer) GetViewGroup() *string {
	if o == nil {
		return nil
	}
	return o.ViewGroup
}

func (o *GetMetadataChildrenMediaContainer) GetViewMode() *int {
	if o == nil {
		return nil
	}
	return o.ViewMode
}

func (o *GetMetadataChildrenMediaContainer) GetDirectory() []GetMetadataChildrenDirectory {
	if o == nil {
		return nil
	}
	return o.Directory
}

func (o *GetMetadataChildrenMediaContainer) GetMetadata() []GetMetadataChildrenMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetMetadataChildrenResponseBody - The children of the library item.
type GetMetadataChildrenResponseBody struct {
	MediaContainer *GetMetadataChildrenMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetMetadataChildrenResponseBody) GetMediaContainer() *GetMetadataChildrenMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetMetadataChildrenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The children of the library item.
	Object *GetMetadataChildrenResponseBody
}

func (o *GetMetadataChildrenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMetadataChildrenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMetadataChildrenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMetadataChildrenResponse) GetObject() *GetMetadataChildrenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
