// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/protocol/security"
)

// LoaderID Unique loader identifier.
type LoaderID string

// RequestID Unique request identifier.
type RequestID string

// InterceptionID Unique intercepted request identifier.
type InterceptionID string

// ErrorReason Network level fetch failure reason.
type ErrorReason int

// ErrorReason as enums.
const (
	ErrorReasonNotSet ErrorReason = iota
	ErrorReasonFailed
	ErrorReasonAborted
	ErrorReasonTimedOut
	ErrorReasonAccessDenied
	ErrorReasonConnectionClosed
	ErrorReasonConnectionReset
	ErrorReasonConnectionRefused
	ErrorReasonConnectionAborted
	ErrorReasonConnectionFailed
	ErrorReasonNameNotResolved
	ErrorReasonInternetDisconnected
	ErrorReasonAddressUnreachable
)

// Valid returns true if enum is set.
func (e ErrorReason) Valid() bool {
	return e >= 1 && e <= 12
}

func (e ErrorReason) String() string {
	switch e {
	case 0:
		return "ErrorReasonNotSet"
	case 1:
		return "Failed"
	case 2:
		return "Aborted"
	case 3:
		return "TimedOut"
	case 4:
		return "AccessDenied"
	case 5:
		return "ConnectionClosed"
	case 6:
		return "ConnectionReset"
	case 7:
		return "ConnectionRefused"
	case 8:
		return "ConnectionAborted"
	case 9:
		return "ConnectionFailed"
	case 10:
		return "NameNotResolved"
	case 11:
		return "InternetDisconnected"
	case 12:
		return "AddressUnreachable"
	}
	return fmt.Sprintf("ErrorReason(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e ErrorReason) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("network.ErrorReason: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *ErrorReason) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"Failed\"":
		*e = 1
	case "\"Aborted\"":
		*e = 2
	case "\"TimedOut\"":
		*e = 3
	case "\"AccessDenied\"":
		*e = 4
	case "\"ConnectionClosed\"":
		*e = 5
	case "\"ConnectionReset\"":
		*e = 6
	case "\"ConnectionRefused\"":
		*e = 7
	case "\"ConnectionAborted\"":
		*e = 8
	case "\"ConnectionFailed\"":
		*e = 9
	case "\"NameNotResolved\"":
		*e = 10
	case "\"InternetDisconnected\"":
		*e = 11
	case "\"AddressUnreachable\"":
		*e = 12
	default:
		return fmt.Errorf("network.ErrorReason: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// Timestamp Number of seconds since epoch.
type Timestamp float64

// String calls (time.Time).String().
func (t Timestamp) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t Timestamp) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("network.Timestamp: " + err.Error())
	}
	*t = Timestamp(f)
	return nil
}

var _ json.Marshaler = (*Timestamp)(nil)
var _ json.Unmarshaler = (*Timestamp)(nil)

// Headers Request / response headers as keys / values of JSON object.
type Headers []byte

// MarshalJSON copies behavior of json.RawMessage.
func (h Headers) MarshalJSON() ([]byte, error) {
	if h == nil {
		return []byte("null"), nil
	}
	return h, nil
}

// UnmarshalJSON copies behavior of json.RawMessage.
func (h *Headers) UnmarshalJSON(data []byte) error {
	if h == nil {
		return errors.New("network.Headers: UnmarshalJSON on nil pointer")
	}
	*h = append((*h)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*Headers)(nil)
var _ json.Unmarshaler = (*Headers)(nil)

// ConnectionType Loading priority of a resource request.
type ConnectionType int

// ConnectionType as enums.
const (
	ConnectionTypeNotSet ConnectionType = iota
	ConnectionTypeNone
	ConnectionTypeCellular2g
	ConnectionTypeCellular3g
	ConnectionTypeCellular4g
	ConnectionTypeBluetooth
	ConnectionTypeEthernet
	ConnectionTypeWifi
	ConnectionTypeWimax
	ConnectionTypeOther
)

// Valid returns true if enum is set.
func (e ConnectionType) Valid() bool {
	return e >= 1 && e <= 9
}

func (e ConnectionType) String() string {
	switch e {
	case 0:
		return "ConnectionTypeNotSet"
	case 1:
		return "none"
	case 2:
		return "cellular2g"
	case 3:
		return "cellular3g"
	case 4:
		return "cellular4g"
	case 5:
		return "bluetooth"
	case 6:
		return "ethernet"
	case 7:
		return "wifi"
	case 8:
		return "wimax"
	case 9:
		return "other"
	}
	return fmt.Sprintf("ConnectionType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e ConnectionType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("network.ConnectionType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *ConnectionType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"none\"":
		*e = 1
	case "\"cellular2g\"":
		*e = 2
	case "\"cellular3g\"":
		*e = 3
	case "\"cellular4g\"":
		*e = 4
	case "\"bluetooth\"":
		*e = 5
	case "\"ethernet\"":
		*e = 6
	case "\"wifi\"":
		*e = 7
	case "\"wimax\"":
		*e = 8
	case "\"other\"":
		*e = 9
	default:
		return fmt.Errorf("network.ConnectionType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// CookieSameSite Represents the cookie's 'SameSite' status: https://tools.ietf.org/html/draft-west-first-party-cookies
type CookieSameSite int

// CookieSameSite as enums.
const (
	CookieSameSiteNotSet CookieSameSite = iota
	CookieSameSiteStrict
	CookieSameSiteLax
)

// Valid returns true if enum is set.
func (e CookieSameSite) Valid() bool {
	return e >= 1 && e <= 2
}

func (e CookieSameSite) String() string {
	switch e {
	case 0:
		return "CookieSameSiteNotSet"
	case 1:
		return "Strict"
	case 2:
		return "Lax"
	}
	return fmt.Sprintf("CookieSameSite(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e CookieSameSite) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("network.CookieSameSite: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *CookieSameSite) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"Strict\"":
		*e = 1
	case "\"Lax\"":
		*e = 2
	default:
		return fmt.Errorf("network.CookieSameSite: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// ResourceTiming Timing information for the request.
type ResourceTiming struct {
	RequestTime       float64 `json:"requestTime"`       // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart        float64 `json:"proxyStart"`        // Started resolving proxy.
	ProxyEnd          float64 `json:"proxyEnd"`          // Finished resolving proxy.
	DNSStart          float64 `json:"dnsStart"`          // Started DNS address resolve.
	DNSEnd            float64 `json:"dnsEnd"`            // Finished DNS address resolve.
	ConnectStart      float64 `json:"connectStart"`      // Started connecting to the remote host.
	ConnectEnd        float64 `json:"connectEnd"`        // Connected to the remote host.
	SslStart          float64 `json:"sslStart"`          // Started SSL handshake.
	SslEnd            float64 `json:"sslEnd"`            // Finished SSL handshake.
	WorkerStart       float64 `json:"workerStart"`       // Started running ServiceWorker.
	WorkerReady       float64 `json:"workerReady"`       // Finished Starting ServiceWorker.
	SendStart         float64 `json:"sendStart"`         // Started sending request.
	SendEnd           float64 `json:"sendEnd"`           // Finished sending request.
	PushStart         float64 `json:"pushStart"`         // Time the server started pushing request.
	PushEnd           float64 `json:"pushEnd"`           // Time the server finished pushing request.
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"` // Finished receiving response headers.
}

// ResourcePriority Loading priority of a resource request.
type ResourcePriority int

// ResourcePriority as enums.
const (
	ResourcePriorityNotSet ResourcePriority = iota
	ResourcePriorityVeryLow
	ResourcePriorityLow
	ResourcePriorityMedium
	ResourcePriorityHigh
	ResourcePriorityVeryHigh
)

// Valid returns true if enum is set.
func (e ResourcePriority) Valid() bool {
	return e >= 1 && e <= 5
}

func (e ResourcePriority) String() string {
	switch e {
	case 0:
		return "ResourcePriorityNotSet"
	case 1:
		return "VeryLow"
	case 2:
		return "Low"
	case 3:
		return "Medium"
	case 4:
		return "High"
	case 5:
		return "VeryHigh"
	}
	return fmt.Sprintf("ResourcePriority(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e ResourcePriority) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("network.ResourcePriority: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *ResourcePriority) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"VeryLow\"":
		*e = 1
	case "\"Low\"":
		*e = 2
	case "\"Medium\"":
		*e = 3
	case "\"High\"":
		*e = 4
	case "\"VeryHigh\"":
		*e = 5
	default:
		return fmt.Errorf("network.ResourcePriority: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// Request HTTP request data.
type Request struct {
	URL              string           `json:"url"`                        // Request URL.
	Method           string           `json:"method"`                     // HTTP request method.
	Headers          Headers          `json:"headers"`                    // HTTP request headers.
	PostData         *string          `json:"postData,omitempty"`         // HTTP POST request data.
	MixedContentType *string          `json:"mixedContentType,omitempty"` // The mixed content status of the request, as defined in http://www.w3.org/TR/mixed-content/
	InitialPriority  ResourcePriority `json:"initialPriority"`            // Priority of the resource request at the time request is sent.
	ReferrerPolicy   string           `json:"referrerPolicy"`             // The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	IsLinkPreload    *bool            `json:"isLinkPreload,omitempty"`    // Whether is loaded via link preload.
}

// SignedCertificateTimestamp Details of a signed certificate timestamp (SCT).
type SignedCertificateTimestamp struct {
	Status             string    `json:"status"`             // Validation status.
	Origin             string    `json:"origin"`             // Origin.
	LogDescription     string    `json:"logDescription"`     // Log name / description.
	LogID              string    `json:"logId"`              // Log ID.
	Timestamp          Timestamp `json:"timestamp"`          // Issuance date.
	HashAlgorithm      string    `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string    `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string    `json:"signatureData"`      // Signature data.
}

// SecurityDetails Security details about a request.
type SecurityDetails struct {
	Protocol                       string                       `json:"protocol"`                       // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                    string                       `json:"keyExchange"`                    // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup               *string                      `json:"keyExchangeGroup,omitempty"`     // (EC)DH group used by the connection, if applicable.
	Cipher                         string                       `json:"cipher"`                         // Cipher name.
	Mac                            *string                      `json:"mac,omitempty"`                  // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateID                  security.CertificateID       `json:"certificateId"`                  // Certificate ID value.
	SubjectName                    string                       `json:"subjectName"`                    // Certificate subject name.
	SanList                        []string                     `json:"sanList"`                        // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                         string                       `json:"issuer"`                         // Name of the issuing CA.
	ValidFrom                      Timestamp                    `json:"validFrom"`                      // Certificate valid from date.
	ValidTo                        Timestamp                    `json:"validTo"`                        // Certificate valid to (expiration) date
	SignedCertificateTimestampList []SignedCertificateTimestamp `json:"signedCertificateTimestampList"` // List of signed certificate timestamps (SCTs).
}

// BlockedReason The reason why request was blocked.
type BlockedReason int

// BlockedReason as enums.
const (
	BlockedReasonNotSet BlockedReason = iota
	BlockedReasonCsp
	BlockedReasonMixedContent
	BlockedReasonOrigin
	BlockedReasonInspector
	BlockedReasonSubresourceFilter
	BlockedReasonOther
)

// Valid returns true if enum is set.
func (e BlockedReason) Valid() bool {
	return e >= 1 && e <= 6
}

func (e BlockedReason) String() string {
	switch e {
	case 0:
		return "BlockedReasonNotSet"
	case 1:
		return "csp"
	case 2:
		return "mixed-content"
	case 3:
		return "origin"
	case 4:
		return "inspector"
	case 5:
		return "subresource-filter"
	case 6:
		return "other"
	}
	return fmt.Sprintf("BlockedReason(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e BlockedReason) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("network.BlockedReason: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *BlockedReason) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"csp\"":
		*e = 1
	case "\"mixed-content\"":
		*e = 2
	case "\"origin\"":
		*e = 3
	case "\"inspector\"":
		*e = 4
	case "\"subresource-filter\"":
		*e = 5
	case "\"other\"":
		*e = 6
	default:
		return fmt.Errorf("network.BlockedReason: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// Response HTTP response data.
type Response struct {
	URL                string           `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             float64          `json:"status"`                       // HTTP response status code.
	StatusText         string           `json:"statusText"`                   // HTTP response status text.
	Headers            Headers          `json:"headers"`                      // HTTP response headers.
	HeadersText        *string          `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string           `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     Headers          `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText *string          `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool             `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionID       float64          `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    *string          `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         *int             `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      *bool            `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  *bool            `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	EncodedDataLength  float64          `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *ResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           *string          `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      security.State   `json:"securityState"`                // Security state of the request resource.
	SecurityDetails    *SecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocketRequest WebSocket request data.
type WebSocketRequest struct {
	Headers Headers `json:"headers"` // HTTP request headers.
}

// WebSocketResponse WebSocket response data.
type WebSocketResponse struct {
	Status             float64 `json:"status"`                       // HTTP response status code.
	StatusText         string  `json:"statusText"`                   // HTTP response status text.
	Headers            Headers `json:"headers"`                      // HTTP response headers.
	HeadersText        *string `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     Headers `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText *string `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocketFrame WebSocket frame data.
type WebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket frame opcode.
	Mask        bool    `json:"mask"`        // WebSocke frame mask.
	PayloadData string  `json:"payloadData"` // WebSocke frame payload data.
}

// Initiator Information about the request initiator.
type Initiator struct {
	Type       string              `json:"type"`                 // Type of this initiator.
	Stack      *runtime.StackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	URL        *string             `json:"url,omitempty"`        // Initiator URL, set for Parser type or for Script type (when script is importing module).
	LineNumber *float64            `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
}

// Cookie Cookie object
type Cookie struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	Domain   string         `json:"domain"`             // Cookie domain.
	Path     string         `json:"path"`               // Cookie path.
	Expires  float64        `json:"expires"`            // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size     int            `json:"size"`               // Cookie size.
	HTTPOnly bool           `json:"httpOnly"`           // True if cookie is http-only.
	Secure   bool           `json:"secure"`             // True if cookie is secure.
	Session  bool           `json:"session"`            // True in case of session cookie.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
}

// AuthChallenge Authorization challenge for HTTP status code 401 or 407.
type AuthChallenge struct {
	Source *string `json:"source,omitempty"` // Source of the authentication challenge.
	Origin string  `json:"origin"`           // Origin of the challenger.
	Scheme string  `json:"scheme"`           // The authentication scheme used, such as basic or digest
	Realm  string  `json:"realm"`            // The realm of the challenge. May be empty.
}

// AuthChallengeResponse Response to an AuthChallenge.
type AuthChallengeResponse struct {
	Response string  `json:"response"`           // The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username *string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password *string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}
