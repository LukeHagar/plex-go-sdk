// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/plexgo/internal/utils"
	"github.com/LukeHagar/plexgo/types"
	"net/http"
)

// GetPlaylistContentsQueryParamType - The type of media to retrieve.
// 1 = movie
// 2 = show
// 3 = season
// 4 = episode
// E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries
type GetPlaylistContentsQueryParamType int64

const (
	GetPlaylistContentsQueryParamTypeMovie   GetPlaylistContentsQueryParamType = 1
	GetPlaylistContentsQueryParamTypeShow    GetPlaylistContentsQueryParamType = 2
	GetPlaylistContentsQueryParamTypeSeason  GetPlaylistContentsQueryParamType = 3
	GetPlaylistContentsQueryParamTypeEpisode GetPlaylistContentsQueryParamType = 4
)

func (e GetPlaylistContentsQueryParamType) ToPointer() *GetPlaylistContentsQueryParamType {
	return &e
}
func (e *GetPlaylistContentsQueryParamType) UnmarshalJSON(data []byte) error {
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
		*e = GetPlaylistContentsQueryParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetPlaylistContentsQueryParamType: %v", v)
	}
}

type GetPlaylistContentsRequest struct {
	// the ID of the playlist
	PlaylistID float64 `pathParam:"style=simple,explode=false,name=playlistID"`
	// The type of media to retrieve.
	// 1 = movie
	// 2 = show
	// 3 = season
	// 4 = episode
	// E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries
	//
	Type GetPlaylistContentsQueryParamType `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetPlaylistContentsRequest) GetPlaylistID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlaylistID
}

func (o *GetPlaylistContentsRequest) GetType() GetPlaylistContentsQueryParamType {
	if o == nil {
		return GetPlaylistContentsQueryParamType(0)
	}
	return o.Type
}

type GetPlaylistContentsPart struct {
	ID                    *int    `json:"id,omitempty"`
	Key                   *string `json:"key,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	File                  *string `json:"file,omitempty"`
	Size                  *int    `json:"size,omitempty"`
	AudioProfile          *string `json:"audioProfile,omitempty"`
	Container             *string `json:"container,omitempty"`
	Has64bitOffsets       *bool   `json:"has64bitOffsets,omitempty"`
	OptimizedForStreaming *bool   `json:"optimizedForStreaming,omitempty"`
	VideoProfile          *string `json:"videoProfile,omitempty"`
}

func (o *GetPlaylistContentsPart) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPlaylistContentsPart) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetPlaylistContentsPart) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetPlaylistContentsPart) GetFile() *string {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *GetPlaylistContentsPart) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetPlaylistContentsPart) GetAudioProfile() *string {
	if o == nil {
		return nil
	}
	return o.AudioProfile
}

func (o *GetPlaylistContentsPart) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetPlaylistContentsPart) GetHas64bitOffsets() *bool {
	if o == nil {
		return nil
	}
	return o.Has64bitOffsets
}

func (o *GetPlaylistContentsPart) GetOptimizedForStreaming() *bool {
	if o == nil {
		return nil
	}
	return o.OptimizedForStreaming
}

func (o *GetPlaylistContentsPart) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

type GetPlaylistContentsMedia struct {
	ID                    *int                      `json:"id,omitempty"`
	Duration              *int                      `json:"duration,omitempty"`
	Bitrate               *int                      `json:"bitrate,omitempty"`
	Width                 *int                      `json:"width,omitempty"`
	Height                *int                      `json:"height,omitempty"`
	AspectRatio           *float64                  `json:"aspectRatio,omitempty"`
	AudioChannels         *int                      `json:"audioChannels,omitempty"`
	AudioCodec            *string                   `json:"audioCodec,omitempty"`
	VideoCodec            *string                   `json:"videoCodec,omitempty"`
	VideoResolution       *string                   `json:"videoResolution,omitempty"`
	Container             *string                   `json:"container,omitempty"`
	VideoFrameRate        *string                   `json:"videoFrameRate,omitempty"`
	OptimizedForStreaming *int                      `json:"optimizedForStreaming,omitempty"`
	AudioProfile          *string                   `json:"audioProfile,omitempty"`
	Has64bitOffsets       *bool                     `json:"has64bitOffsets,omitempty"`
	VideoProfile          *string                   `json:"videoProfile,omitempty"`
	Part                  []GetPlaylistContentsPart `json:"Part,omitempty"`
}

func (o *GetPlaylistContentsMedia) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPlaylistContentsMedia) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetPlaylistContentsMedia) GetBitrate() *int {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *GetPlaylistContentsMedia) GetWidth() *int {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *GetPlaylistContentsMedia) GetHeight() *int {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *GetPlaylistContentsMedia) GetAspectRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.AspectRatio
}

func (o *GetPlaylistContentsMedia) GetAudioChannels() *int {
	if o == nil {
		return nil
	}
	return o.AudioChannels
}

func (o *GetPlaylistContentsMedia) GetAudioCodec() *string {
	if o == nil {
		return nil
	}
	return o.AudioCodec
}

func (o *GetPlaylistContentsMedia) GetVideoCodec() *string {
	if o == nil {
		return nil
	}
	return o.VideoCodec
}

func (o *GetPlaylistContentsMedia) GetVideoResolution() *string {
	if o == nil {
		return nil
	}
	return o.VideoResolution
}

func (o *GetPlaylistContentsMedia) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetPlaylistContentsMedia) GetVideoFrameRate() *string {
	if o == nil {
		return nil
	}
	return o.VideoFrameRate
}

func (o *GetPlaylistContentsMedia) GetOptimizedForStreaming() *int {
	if o == nil {
		return nil
	}
	return o.OptimizedForStreaming
}

func (o *GetPlaylistContentsMedia) GetAudioProfile() *string {
	if o == nil {
		return nil
	}
	return o.AudioProfile
}

func (o *GetPlaylistContentsMedia) GetHas64bitOffsets() *bool {
	if o == nil {
		return nil
	}
	return o.Has64bitOffsets
}

func (o *GetPlaylistContentsMedia) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

func (o *GetPlaylistContentsMedia) GetPart() []GetPlaylistContentsPart {
	if o == nil {
		return nil
	}
	return o.Part
}

type GetPlaylistContentsGenre struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetPlaylistContentsGenre) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetPlaylistContentsCountry struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetPlaylistContentsCountry) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetPlaylistContentsDirector struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetPlaylistContentsDirector) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetPlaylistContentsWriter struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetPlaylistContentsWriter) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetPlaylistContentsRole struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetPlaylistContentsRole) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetPlaylistContentsMetadata struct {
	RatingKey              *string                       `json:"ratingKey,omitempty"`
	Key                    *string                       `json:"key,omitempty"`
	GUID                   *string                       `json:"guid,omitempty"`
	Studio                 *string                       `json:"studio,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	Title                  *string                       `json:"title,omitempty"`
	TitleSort              *string                       `json:"titleSort,omitempty"`
	LibrarySectionTitle    *string                       `json:"librarySectionTitle,omitempty"`
	LibrarySectionID       *int                          `json:"librarySectionID,omitempty"`
	LibrarySectionKey      *string                       `json:"librarySectionKey,omitempty"`
	ContentRating          *string                       `json:"contentRating,omitempty"`
	Summary                *string                       `json:"summary,omitempty"`
	Rating                 *float64                      `json:"rating,omitempty"`
	AudienceRating         *float64                      `json:"audienceRating,omitempty"`
	Year                   *int                          `json:"year,omitempty"`
	Tagline                *string                       `json:"tagline,omitempty"`
	Thumb                  *string                       `json:"thumb,omitempty"`
	Art                    *string                       `json:"art,omitempty"`
	Duration               *int                          `json:"duration,omitempty"`
	OriginallyAvailableAt  *types.Date                   `json:"originallyAvailableAt,omitempty"`
	AddedAt                *int                          `json:"addedAt,omitempty"`
	UpdatedAt              *int                          `json:"updatedAt,omitempty"`
	AudienceRatingImage    *string                       `json:"audienceRatingImage,omitempty"`
	HasPremiumExtras       *string                       `json:"hasPremiumExtras,omitempty"`
	HasPremiumPrimaryExtra *string                       `json:"hasPremiumPrimaryExtra,omitempty"`
	RatingImage            *string                       `json:"ratingImage,omitempty"`
	Media                  []GetPlaylistContentsMedia    `json:"Media,omitempty"`
	Genre                  []GetPlaylistContentsGenre    `json:"Genre,omitempty"`
	Country                []GetPlaylistContentsCountry  `json:"Country,omitempty"`
	Director               []GetPlaylistContentsDirector `json:"Director,omitempty"`
	Writer                 []GetPlaylistContentsWriter   `json:"Writer,omitempty"`
	Role                   []GetPlaylistContentsRole     `json:"Role,omitempty"`
}

func (g GetPlaylistContentsMetadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPlaylistContentsMetadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPlaylistContentsMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetPlaylistContentsMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetPlaylistContentsMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetPlaylistContentsMetadata) GetStudio() *string {
	if o == nil {
		return nil
	}
	return o.Studio
}

func (o *GetPlaylistContentsMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetPlaylistContentsMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetPlaylistContentsMetadata) GetTitleSort() *string {
	if o == nil {
		return nil
	}
	return o.TitleSort
}

func (o *GetPlaylistContentsMetadata) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetPlaylistContentsMetadata) GetLibrarySectionID() *int {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetPlaylistContentsMetadata) GetLibrarySectionKey() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionKey
}

func (o *GetPlaylistContentsMetadata) GetContentRating() *string {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetPlaylistContentsMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetPlaylistContentsMetadata) GetRating() *float64 {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GetPlaylistContentsMetadata) GetAudienceRating() *float64 {
	if o == nil {
		return nil
	}
	return o.AudienceRating
}

func (o *GetPlaylistContentsMetadata) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *GetPlaylistContentsMetadata) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *GetPlaylistContentsMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetPlaylistContentsMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetPlaylistContentsMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetPlaylistContentsMetadata) GetOriginallyAvailableAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.OriginallyAvailableAt
}

func (o *GetPlaylistContentsMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetPlaylistContentsMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetPlaylistContentsMetadata) GetAudienceRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.AudienceRatingImage
}

func (o *GetPlaylistContentsMetadata) GetHasPremiumExtras() *string {
	if o == nil {
		return nil
	}
	return o.HasPremiumExtras
}

func (o *GetPlaylistContentsMetadata) GetHasPremiumPrimaryExtra() *string {
	if o == nil {
		return nil
	}
	return o.HasPremiumPrimaryExtra
}

func (o *GetPlaylistContentsMetadata) GetRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.RatingImage
}

func (o *GetPlaylistContentsMetadata) GetMedia() []GetPlaylistContentsMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *GetPlaylistContentsMetadata) GetGenre() []GetPlaylistContentsGenre {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetPlaylistContentsMetadata) GetCountry() []GetPlaylistContentsCountry {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetPlaylistContentsMetadata) GetDirector() []GetPlaylistContentsDirector {
	if o == nil {
		return nil
	}
	return o.Director
}

func (o *GetPlaylistContentsMetadata) GetWriter() []GetPlaylistContentsWriter {
	if o == nil {
		return nil
	}
	return o.Writer
}

func (o *GetPlaylistContentsMetadata) GetRole() []GetPlaylistContentsRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type GetPlaylistContentsMediaContainer struct {
	Size         *int                          `json:"size,omitempty"`
	Composite    *string                       `json:"composite,omitempty"`
	Duration     *int                          `json:"duration,omitempty"`
	LeafCount    *int                          `json:"leafCount,omitempty"`
	PlaylistType *string                       `json:"playlistType,omitempty"`
	RatingKey    *string                       `json:"ratingKey,omitempty"`
	Smart        *bool                         `json:"smart,omitempty"`
	Title        *string                       `json:"title,omitempty"`
	Metadata     []GetPlaylistContentsMetadata `json:"Metadata,omitempty"`
}

func (o *GetPlaylistContentsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetPlaylistContentsMediaContainer) GetComposite() *string {
	if o == nil {
		return nil
	}
	return o.Composite
}

func (o *GetPlaylistContentsMediaContainer) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetPlaylistContentsMediaContainer) GetLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *GetPlaylistContentsMediaContainer) GetPlaylistType() *string {
	if o == nil {
		return nil
	}
	return o.PlaylistType
}

func (o *GetPlaylistContentsMediaContainer) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetPlaylistContentsMediaContainer) GetSmart() *bool {
	if o == nil {
		return nil
	}
	return o.Smart
}

func (o *GetPlaylistContentsMediaContainer) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetPlaylistContentsMediaContainer) GetMetadata() []GetPlaylistContentsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetPlaylistContentsResponseBody - The playlist contents
type GetPlaylistContentsResponseBody struct {
	MediaContainer *GetPlaylistContentsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetPlaylistContentsResponseBody) GetMediaContainer() *GetPlaylistContentsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetPlaylistContentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The playlist contents
	Object *GetPlaylistContentsResponseBody
}

func (o *GetPlaylistContentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPlaylistContentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPlaylistContentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPlaylistContentsResponse) GetObject() *GetPlaylistContentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
