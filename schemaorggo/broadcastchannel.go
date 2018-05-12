package schemaorggo

import "encoding/json"

// BroadcastChannel see : https://schema.org/BroadcastChannel
type BroadcastChannel struct {
	Intangible

	typeContext

	// BroadcastChannelId see : https://schema.org/broadcastChannelId
	// The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
	// types : Text
	BroadcastChannelId string `json:"broadcastChannelId,omitempty"`

	// BroadcastFrequency see : http://pending.schema.org/broadcastFrequency
	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	// types : BroadcastFrequencySpecification Text
	BroadcastFrequency interface{} `json:"broadcastFrequency,omitempty"`

	// BroadcastServiceTier see : https://schema.org/broadcastServiceTier
	// The type of service required to have access to the channel (e.g. Standard or Premium).
	// types : Text
	BroadcastServiceTier string `json:"broadcastServiceTier,omitempty"`

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	// types : Text URL
	Genre string `json:"genre,omitempty"`

	// InBroadcastLineup see : https://schema.org/inBroadcastLineup
	// The CableOrSatelliteService offering the channel.
	// types : CableOrSatelliteService
	InBroadcastLineup *CableOrSatelliteService `json:"inBroadcastLineup,omitempty"`

	// ProvidesBroadcastService see : https://schema.org/providesBroadcastService
	// The BroadcastService offered on this channel. Inverse property: hasBroadcastChannel (see: https://schema.orghttp://pending.schema.org/hasBroadcastChannel).
	// types : BroadcastService
	ProvidesBroadcastService *BroadcastService `json:"providesBroadcastService,omitempty"`
}

func (v BroadcastChannel) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BroadcastChannel"

	return json.Marshal(v)
}

func (v *BroadcastChannel) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BroadcastChannel"

	return json.Marshal(*v)
}
