// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/plexgo/internal/utils"
	"net/http"
	"time"
)

var GetServerResourcesServerList = []string{
	"https://plex.tv/api/v2/",
}

type GetServerResourcesGlobals struct {
	// The unique identifier for the client application
	// This is used to track the client application and its usage
	// (UUID, serial number, or other number unique per device)
	//
	ClientID *string `queryParam:"style=form,explode=true,name=X-Plex-Client-Identifier"`
}

func (o *GetServerResourcesGlobals) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

// IncludeHTTPS - Include Https entries in the results
type IncludeHTTPS int

const (
	IncludeHTTPSZero IncludeHTTPS = 0
	IncludeHTTPSOne  IncludeHTTPS = 1
)

func (e IncludeHTTPS) ToPointer() *IncludeHTTPS {
	return &e
}
func (e *IncludeHTTPS) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = IncludeHTTPS(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncludeHTTPS: %v", v)
	}
}

// IncludeRelay - Include Relay addresses in the results
// E.g: https://10-0-0-25.bbf8e10c7fa20447cacee74cd9914cde.plex.direct:32400
type IncludeRelay int

const (
	IncludeRelayZero IncludeRelay = 0
	IncludeRelayOne  IncludeRelay = 1
)

func (e IncludeRelay) ToPointer() *IncludeRelay {
	return &e
}
func (e *IncludeRelay) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = IncludeRelay(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncludeRelay: %v", v)
	}
}

// IncludeIPv6 - Include IPv6 entries in the results
type IncludeIPv6 int

const (
	IncludeIPv6Zero IncludeIPv6 = 0
	IncludeIPv6One  IncludeIPv6 = 1
)

func (e IncludeIPv6) ToPointer() *IncludeIPv6 {
	return &e
}
func (e *IncludeIPv6) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = IncludeIPv6(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncludeIPv6: %v", v)
	}
}

type GetServerResourcesRequest struct {
	// The unique identifier for the client application
	// This is used to track the client application and its usage
	// (UUID, serial number, or other number unique per device)
	//
	ClientID *string `queryParam:"style=form,explode=true,name=X-Plex-Client-Identifier"`
	// Include Https entries in the results
	IncludeHTTPS *IncludeHTTPS `default:"0" queryParam:"style=form,explode=true,name=includeHttps"`
	// Include Relay addresses in the results
	// E.g: https://10-0-0-25.bbf8e10c7fa20447cacee74cd9914cde.plex.direct:32400
	//
	IncludeRelay *IncludeRelay `default:"0" queryParam:"style=form,explode=true,name=includeRelay"`
	// Include IPv6 entries in the results
	IncludeIPv6 *IncludeIPv6 `default:"0" queryParam:"style=form,explode=true,name=includeIPv6"`
}

func (g GetServerResourcesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetServerResourcesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetServerResourcesRequest) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *GetServerResourcesRequest) GetIncludeHTTPS() *IncludeHTTPS {
	if o == nil {
		return nil
	}
	return o.IncludeHTTPS
}

func (o *GetServerResourcesRequest) GetIncludeRelay() *IncludeRelay {
	if o == nil {
		return nil
	}
	return o.IncludeRelay
}

func (o *GetServerResourcesRequest) GetIncludeIPv6() *IncludeIPv6 {
	if o == nil {
		return nil
	}
	return o.IncludeIPv6
}

type Connections struct {
	Protocol string  `json:"protocol"`
	Address  string  `json:"address"`
	Port     float64 `json:"port"`
	URI      string  `json:"uri"`
	Local    bool    `json:"local"`
	Relay    bool    `json:"relay"`
	IPv6     bool    `json:"IPv6"`
}

func (o *Connections) GetProtocol() string {
	if o == nil {
		return ""
	}
	return o.Protocol
}

func (o *Connections) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *Connections) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *Connections) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

func (o *Connections) GetLocal() bool {
	if o == nil {
		return false
	}
	return o.Local
}

func (o *Connections) GetRelay() bool {
	if o == nil {
		return false
	}
	return o.Relay
}

func (o *Connections) GetIPv6() bool {
	if o == nil {
		return false
	}
	return o.IPv6
}

type PlexDevice struct {
	Name             string    `json:"name"`
	Product          string    `json:"product"`
	ProductVersion   string    `json:"productVersion"`
	Platform         *string   `json:"platform"`
	PlatformVersion  *string   `json:"platformVersion"`
	Device           *string   `json:"device"`
	ClientIdentifier string    `json:"clientIdentifier"`
	CreatedAt        time.Time `json:"createdAt"`
	LastSeenAt       time.Time `json:"lastSeenAt"`
	Provides         string    `json:"provides"`
	// ownerId is null when the device is owned by the token used to send the request
	OwnerID                *int64        `json:"ownerId"`
	SourceTitle            *string       `json:"sourceTitle"`
	PublicAddress          string        `json:"publicAddress"`
	AccessToken            string        `json:"accessToken"`
	Owned                  bool          `json:"owned"`
	Home                   bool          `json:"home"`
	Synced                 bool          `json:"synced"`
	Relay                  bool          `json:"relay"`
	Presence               bool          `json:"presence"`
	HTTPSRequired          bool          `json:"httpsRequired"`
	PublicAddressMatches   bool          `json:"publicAddressMatches"`
	DNSRebindingProtection bool          `json:"dnsRebindingProtection"`
	NatLoopbackSupported   bool          `json:"natLoopbackSupported"`
	Connections            []Connections `json:"connections"`
}

func (p PlexDevice) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PlexDevice) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PlexDevice) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PlexDevice) GetProduct() string {
	if o == nil {
		return ""
	}
	return o.Product
}

func (o *PlexDevice) GetProductVersion() string {
	if o == nil {
		return ""
	}
	return o.ProductVersion
}

func (o *PlexDevice) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *PlexDevice) GetPlatformVersion() *string {
	if o == nil {
		return nil
	}
	return o.PlatformVersion
}

func (o *PlexDevice) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *PlexDevice) GetClientIdentifier() string {
	if o == nil {
		return ""
	}
	return o.ClientIdentifier
}

func (o *PlexDevice) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PlexDevice) GetLastSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastSeenAt
}

func (o *PlexDevice) GetProvides() string {
	if o == nil {
		return ""
	}
	return o.Provides
}

func (o *PlexDevice) GetOwnerID() *int64 {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *PlexDevice) GetSourceTitle() *string {
	if o == nil {
		return nil
	}
	return o.SourceTitle
}

func (o *PlexDevice) GetPublicAddress() string {
	if o == nil {
		return ""
	}
	return o.PublicAddress
}

func (o *PlexDevice) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *PlexDevice) GetOwned() bool {
	if o == nil {
		return false
	}
	return o.Owned
}

func (o *PlexDevice) GetHome() bool {
	if o == nil {
		return false
	}
	return o.Home
}

func (o *PlexDevice) GetSynced() bool {
	if o == nil {
		return false
	}
	return o.Synced
}

func (o *PlexDevice) GetRelay() bool {
	if o == nil {
		return false
	}
	return o.Relay
}

func (o *PlexDevice) GetPresence() bool {
	if o == nil {
		return false
	}
	return o.Presence
}

func (o *PlexDevice) GetHTTPSRequired() bool {
	if o == nil {
		return false
	}
	return o.HTTPSRequired
}

func (o *PlexDevice) GetPublicAddressMatches() bool {
	if o == nil {
		return false
	}
	return o.PublicAddressMatches
}

func (o *PlexDevice) GetDNSRebindingProtection() bool {
	if o == nil {
		return false
	}
	return o.DNSRebindingProtection
}

func (o *PlexDevice) GetNatLoopbackSupported() bool {
	if o == nil {
		return false
	}
	return o.NatLoopbackSupported
}

func (o *PlexDevice) GetConnections() []Connections {
	if o == nil {
		return []Connections{}
	}
	return o.Connections
}

type GetServerResourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of Plex Devices. This includes Plex hosted servers and clients
	PlexDevices []PlexDevice
}

func (o *GetServerResourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerResourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerResourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServerResourcesResponse) GetPlexDevices() []PlexDevice {
	if o == nil {
		return nil
	}
	return o.PlexDevices
}
