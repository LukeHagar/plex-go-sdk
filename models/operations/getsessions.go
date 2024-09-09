// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetSessionsStream struct {
	AlbumGain            *string `json:"albumGain,omitempty"`
	AlbumPeak            *string `json:"albumPeak,omitempty"`
	AlbumRange           *string `json:"albumRange,omitempty"`
	AudioChannelLayout   *string `json:"audioChannelLayout,omitempty"`
	BitDepth             *int    `json:"bitDepth,omitempty"`
	Bitrate              *int    `json:"bitrate,omitempty"`
	Channels             *int    `json:"channels,omitempty"`
	Codec                *string `json:"codec,omitempty"`
	DisplayTitle         *string `json:"displayTitle,omitempty"`
	ExtendedDisplayTitle *string `json:"extendedDisplayTitle,omitempty"`
	Gain                 *string `json:"gain,omitempty"`
	ID                   *string `json:"id,omitempty"`
	Index                *int    `json:"index,omitempty"`
	Loudness             *string `json:"loudness,omitempty"`
	Lra                  *string `json:"lra,omitempty"`
	Peak                 *string `json:"peak,omitempty"`
	SamplingRate         *int    `json:"samplingRate,omitempty"`
	Selected             *bool   `json:"selected,omitempty"`
	StreamType           *int    `json:"streamType,omitempty"`
	Location             *string `json:"location,omitempty"`
}

func (o *GetSessionsStream) GetAlbumGain() *string {
	if o == nil {
		return nil
	}
	return o.AlbumGain
}

func (o *GetSessionsStream) GetAlbumPeak() *string {
	if o == nil {
		return nil
	}
	return o.AlbumPeak
}

func (o *GetSessionsStream) GetAlbumRange() *string {
	if o == nil {
		return nil
	}
	return o.AlbumRange
}

func (o *GetSessionsStream) GetAudioChannelLayout() *string {
	if o == nil {
		return nil
	}
	return o.AudioChannelLayout
}

func (o *GetSessionsStream) GetBitDepth() *int {
	if o == nil {
		return nil
	}
	return o.BitDepth
}

func (o *GetSessionsStream) GetBitrate() *int {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *GetSessionsStream) GetChannels() *int {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *GetSessionsStream) GetCodec() *string {
	if o == nil {
		return nil
	}
	return o.Codec
}

func (o *GetSessionsStream) GetDisplayTitle() *string {
	if o == nil {
		return nil
	}
	return o.DisplayTitle
}

func (o *GetSessionsStream) GetExtendedDisplayTitle() *string {
	if o == nil {
		return nil
	}
	return o.ExtendedDisplayTitle
}

func (o *GetSessionsStream) GetGain() *string {
	if o == nil {
		return nil
	}
	return o.Gain
}

func (o *GetSessionsStream) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSessionsStream) GetIndex() *int {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetSessionsStream) GetLoudness() *string {
	if o == nil {
		return nil
	}
	return o.Loudness
}

func (o *GetSessionsStream) GetLra() *string {
	if o == nil {
		return nil
	}
	return o.Lra
}

func (o *GetSessionsStream) GetPeak() *string {
	if o == nil {
		return nil
	}
	return o.Peak
}

func (o *GetSessionsStream) GetSamplingRate() *int {
	if o == nil {
		return nil
	}
	return o.SamplingRate
}

func (o *GetSessionsStream) GetSelected() *bool {
	if o == nil {
		return nil
	}
	return o.Selected
}

func (o *GetSessionsStream) GetStreamType() *int {
	if o == nil {
		return nil
	}
	return o.StreamType
}

func (o *GetSessionsStream) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

type GetSessionsPart struct {
	Container    *string             `json:"container,omitempty"`
	Duration     *int                `json:"duration,omitempty"`
	File         *string             `json:"file,omitempty"`
	HasThumbnail *string             `json:"hasThumbnail,omitempty"`
	ID           *string             `json:"id,omitempty"`
	Key          *string             `json:"key,omitempty"`
	Size         *int                `json:"size,omitempty"`
	Decision     *string             `json:"decision,omitempty"`
	Selected     *bool               `json:"selected,omitempty"`
	Stream       []GetSessionsStream `json:"Stream,omitempty"`
}

func (o *GetSessionsPart) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetSessionsPart) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetSessionsPart) GetFile() *string {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *GetSessionsPart) GetHasThumbnail() *string {
	if o == nil {
		return nil
	}
	return o.HasThumbnail
}

func (o *GetSessionsPart) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSessionsPart) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetSessionsPart) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetSessionsPart) GetDecision() *string {
	if o == nil {
		return nil
	}
	return o.Decision
}

func (o *GetSessionsPart) GetSelected() *bool {
	if o == nil {
		return nil
	}
	return o.Selected
}

func (o *GetSessionsPart) GetStream() []GetSessionsStream {
	if o == nil {
		return nil
	}
	return o.Stream
}

type GetSessionsMedia struct {
	AudioChannels *int              `json:"audioChannels,omitempty"`
	AudioCodec    *string           `json:"audioCodec,omitempty"`
	Bitrate       *int              `json:"bitrate,omitempty"`
	Container     *string           `json:"container,omitempty"`
	Duration      *int              `json:"duration,omitempty"`
	ID            *string           `json:"id,omitempty"`
	Selected      *bool             `json:"selected,omitempty"`
	Part          []GetSessionsPart `json:"Part,omitempty"`
}

func (o *GetSessionsMedia) GetAudioChannels() *int {
	if o == nil {
		return nil
	}
	return o.AudioChannels
}

func (o *GetSessionsMedia) GetAudioCodec() *string {
	if o == nil {
		return nil
	}
	return o.AudioCodec
}

func (o *GetSessionsMedia) GetBitrate() *int {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *GetSessionsMedia) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetSessionsMedia) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetSessionsMedia) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSessionsMedia) GetSelected() *bool {
	if o == nil {
		return nil
	}
	return o.Selected
}

func (o *GetSessionsMedia) GetPart() []GetSessionsPart {
	if o == nil {
		return nil
	}
	return o.Part
}

type GetSessionsUser struct {
	ID    *string `json:"id,omitempty"`
	Thumb *string `json:"thumb,omitempty"`
	Title *string `json:"title,omitempty"`
}

func (o *GetSessionsUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSessionsUser) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetSessionsUser) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type Player struct {
	Address             *string `json:"address,omitempty"`
	MachineIdentifier   *string `json:"machineIdentifier,omitempty"`
	Model               *string `json:"model,omitempty"`
	Platform            *string `json:"platform,omitempty"`
	PlatformVersion     *string `json:"platformVersion,omitempty"`
	Product             *string `json:"product,omitempty"`
	Profile             *string `json:"profile,omitempty"`
	RemotePublicAddress *string `json:"remotePublicAddress,omitempty"`
	State               *string `json:"state,omitempty"`
	Title               *string `json:"title,omitempty"`
	Version             *string `json:"version,omitempty"`
	Local               *bool   `json:"local,omitempty"`
	Relayed             *bool   `json:"relayed,omitempty"`
	Secure              *bool   `json:"secure,omitempty"`
	UserID              *int    `json:"userID,omitempty"`
}

func (o *Player) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Player) GetMachineIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.MachineIdentifier
}

func (o *Player) GetModel() *string {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *Player) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *Player) GetPlatformVersion() *string {
	if o == nil {
		return nil
	}
	return o.PlatformVersion
}

func (o *Player) GetProduct() *string {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *Player) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *Player) GetRemotePublicAddress() *string {
	if o == nil {
		return nil
	}
	return o.RemotePublicAddress
}

func (o *Player) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Player) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Player) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *Player) GetLocal() *bool {
	if o == nil {
		return nil
	}
	return o.Local
}

func (o *Player) GetRelayed() *bool {
	if o == nil {
		return nil
	}
	return o.Relayed
}

func (o *Player) GetSecure() *bool {
	if o == nil {
		return nil
	}
	return o.Secure
}

func (o *Player) GetUserID() *int {
	if o == nil {
		return nil
	}
	return o.UserID
}

type Session struct {
	ID        *string `json:"id,omitempty"`
	Bandwidth *int    `json:"bandwidth,omitempty"`
	Location  *string `json:"location,omitempty"`
}

func (o *Session) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Session) GetBandwidth() *int {
	if o == nil {
		return nil
	}
	return o.Bandwidth
}

func (o *Session) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

type GetSessionsMetadata struct {
	AddedAt              *int               `json:"addedAt,omitempty"`
	Art                  *string            `json:"art,omitempty"`
	Duration             *int               `json:"duration,omitempty"`
	GrandparentArt       *string            `json:"grandparentArt,omitempty"`
	GrandparentGUID      *string            `json:"grandparentGuid,omitempty"`
	GrandparentKey       *string            `json:"grandparentKey,omitempty"`
	GrandparentRatingKey *string            `json:"grandparentRatingKey,omitempty"`
	GrandparentThumb     *string            `json:"grandparentThumb,omitempty"`
	GrandparentTitle     *string            `json:"grandparentTitle,omitempty"`
	GUID                 *string            `json:"guid,omitempty"`
	Index                *int               `json:"index,omitempty"`
	Key                  *string            `json:"key,omitempty"`
	LibrarySectionID     *string            `json:"librarySectionID,omitempty"`
	LibrarySectionKey    *string            `json:"librarySectionKey,omitempty"`
	LibrarySectionTitle  *string            `json:"librarySectionTitle,omitempty"`
	MusicAnalysisVersion *string            `json:"musicAnalysisVersion,omitempty"`
	ParentGUID           *string            `json:"parentGuid,omitempty"`
	ParentIndex          *int               `json:"parentIndex,omitempty"`
	ParentKey            *string            `json:"parentKey,omitempty"`
	ParentRatingKey      *string            `json:"parentRatingKey,omitempty"`
	ParentStudio         *string            `json:"parentStudio,omitempty"`
	ParentThumb          *string            `json:"parentThumb,omitempty"`
	ParentTitle          *string            `json:"parentTitle,omitempty"`
	ParentYear           *int               `json:"parentYear,omitempty"`
	RatingCount          *int               `json:"ratingCount,omitempty"`
	RatingKey            *string            `json:"ratingKey,omitempty"`
	SessionKey           *string            `json:"sessionKey,omitempty"`
	Thumb                *string            `json:"thumb,omitempty"`
	Title                *string            `json:"title,omitempty"`
	TitleSort            *string            `json:"titleSort,omitempty"`
	Type                 *string            `json:"type,omitempty"`
	UpdatedAt            *int               `json:"updatedAt,omitempty"`
	ViewOffset           *int               `json:"viewOffset,omitempty"`
	Media                []GetSessionsMedia `json:"Media,omitempty"`
	User                 *GetSessionsUser   `json:"User,omitempty"`
	Player               *Player            `json:"Player,omitempty"`
	Session              *Session           `json:"Session,omitempty"`
}

func (o *GetSessionsMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetSessionsMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetSessionsMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetSessionsMetadata) GetGrandparentArt() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentArt
}

func (o *GetSessionsMetadata) GetGrandparentGUID() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentGUID
}

func (o *GetSessionsMetadata) GetGrandparentKey() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentKey
}

func (o *GetSessionsMetadata) GetGrandparentRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentRatingKey
}

func (o *GetSessionsMetadata) GetGrandparentThumb() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentThumb
}

func (o *GetSessionsMetadata) GetGrandparentTitle() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentTitle
}

func (o *GetSessionsMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetSessionsMetadata) GetIndex() *int {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetSessionsMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetSessionsMetadata) GetLibrarySectionID() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetSessionsMetadata) GetLibrarySectionKey() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionKey
}

func (o *GetSessionsMetadata) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetSessionsMetadata) GetMusicAnalysisVersion() *string {
	if o == nil {
		return nil
	}
	return o.MusicAnalysisVersion
}

func (o *GetSessionsMetadata) GetParentGUID() *string {
	if o == nil {
		return nil
	}
	return o.ParentGUID
}

func (o *GetSessionsMetadata) GetParentIndex() *int {
	if o == nil {
		return nil
	}
	return o.ParentIndex
}

func (o *GetSessionsMetadata) GetParentKey() *string {
	if o == nil {
		return nil
	}
	return o.ParentKey
}

func (o *GetSessionsMetadata) GetParentRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.ParentRatingKey
}

func (o *GetSessionsMetadata) GetParentStudio() *string {
	if o == nil {
		return nil
	}
	return o.ParentStudio
}

func (o *GetSessionsMetadata) GetParentThumb() *string {
	if o == nil {
		return nil
	}
	return o.ParentThumb
}

func (o *GetSessionsMetadata) GetParentTitle() *string {
	if o == nil {
		return nil
	}
	return o.ParentTitle
}

func (o *GetSessionsMetadata) GetParentYear() *int {
	if o == nil {
		return nil
	}
	return o.ParentYear
}

func (o *GetSessionsMetadata) GetRatingCount() *int {
	if o == nil {
		return nil
	}
	return o.RatingCount
}

func (o *GetSessionsMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetSessionsMetadata) GetSessionKey() *string {
	if o == nil {
		return nil
	}
	return o.SessionKey
}

func (o *GetSessionsMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetSessionsMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetSessionsMetadata) GetTitleSort() *string {
	if o == nil {
		return nil
	}
	return o.TitleSort
}

func (o *GetSessionsMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetSessionsMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetSessionsMetadata) GetViewOffset() *int {
	if o == nil {
		return nil
	}
	return o.ViewOffset
}

func (o *GetSessionsMetadata) GetMedia() []GetSessionsMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *GetSessionsMetadata) GetUser() *GetSessionsUser {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *GetSessionsMetadata) GetPlayer() *Player {
	if o == nil {
		return nil
	}
	return o.Player
}

func (o *GetSessionsMetadata) GetSession() *Session {
	if o == nil {
		return nil
	}
	return o.Session
}

type GetSessionsMediaContainer struct {
	Size     *int                  `json:"size,omitempty"`
	Metadata []GetSessionsMetadata `json:"Metadata,omitempty"`
}

func (o *GetSessionsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetSessionsMediaContainer) GetMetadata() []GetSessionsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetSessionsResponseBody - List of Active Plex Sessions
type GetSessionsResponseBody struct {
	MediaContainer *GetSessionsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetSessionsResponseBody) GetMediaContainer() *GetSessionsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetSessionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of Active Plex Sessions
	Object *GetSessionsResponseBody
}

func (o *GetSessionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSessionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSessionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSessionsResponse) GetObject() *GetSessionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
