package schemaorggo

import "encoding/json"

// BroadcastService see : https://schema.org/BroadcastService
type BroadcastService struct {
	Service

	typeContext

	// BroadcastAffiliateOf see : https://schema.org/broadcastAffiliateOf
	// The media network(s) whose content is broadcast on this station.
	// types : Organization
	BroadcastAffiliateOf *Organization `json:"broadcastAffiliateOf,omitempty"`

	// BroadcastDisplayName see : https://schema.org/broadcastDisplayName
	// The name displayed in the channel guide. For many US affiliates, it is the network name.
	// types : Text
	BroadcastDisplayName string `json:"broadcastDisplayName,omitempty"`

	// BroadcastFrequency see : http://pending.schema.org/broadcastFrequency
	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	// types : BroadcastFrequencySpecification Text
	BroadcastFrequency interface{} `json:"broadcastFrequency,omitempty"`

	// BroadcastTimezone see : https://schema.org/broadcastTimezone
	// The timezone in ISO 8601 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601) for which the service bases its broadcasts
	// types : Text
	BroadcastTimezone string `json:"broadcastTimezone,omitempty"`

	// Broadcaster see : https://schema.org/broadcaster
	// The organization owning or operating the broadcast service.
	// types : Organization
	Broadcaster *Organization `json:"broadcaster,omitempty"`

	// HasBroadcastChannel see : http://pending.schema.org/hasBroadcastChannel
	// A broadcast channel of a broadcast service. Inverse property: providesBroadcastService (see: https://schema.org/providesBroadcastService).
	// types : BroadcastChannel
	HasBroadcastChannel *BroadcastChannel `json:"hasBroadcastChannel,omitempty"`

	// ParentService see : https://schema.org/parentService
	// A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	// types : BroadcastService
	ParentService *BroadcastService `json:"parentService,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat string `json:"videoFormat,omitempty"`
}

func (v BroadcastService) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BroadcastService"

	return json.Marshal(v)
}

func (v *BroadcastService) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BroadcastService"

	return json.Marshal(*v)
}
