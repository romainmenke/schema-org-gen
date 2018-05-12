package schemaorggo

import "encoding/json"

// TelevisionChannel see : https://schema.org/TelevisionChannel
type TelevisionChannel struct {
	BroadcastChannel

	typeContext

	// BroadcastChannelId see : https://schema.org/broadcastChannelId
	// The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
	BroadcastChannelId string `json:"broadcastChannelId,omitempty"` // types : Text

	// BroadcastFrequency see : http://pending.schema.org/broadcastFrequency
	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. "87 FM".
	BroadcastFrequency interface{} `json:"broadcastFrequency,omitempty"` // types : BroadcastFrequencySpecification Text

	// BroadcastServiceTier see : https://schema.org/broadcastServiceTier
	// The type of service required to have access to the channel (e.g. Standard or Premium).
	BroadcastServiceTier string `json:"broadcastServiceTier,omitempty"` // types : Text

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	Genre string `json:"genre,omitempty"` // types : Text URL

	// InBroadcastLineup see : https://schema.org/inBroadcastLineup
	// The CableOrSatelliteService offering the channel.
	InBroadcastLineup *CableOrSatelliteService `json:"inBroadcastLineup,omitempty"` // types : CableOrSatelliteService

	// ProvidesBroadcastService see : https://schema.org/providesBroadcastService
	// The BroadcastService offered on this channel. Inverse property: hasBroadcastChannel (see: https://schema.orghttp://pending.schema.org/hasBroadcastChannel).
	ProvidesBroadcastService *BroadcastService `json:"providesBroadcastService,omitempty"` // types : BroadcastService

}

func (v TelevisionChannel) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TelevisionChannel"

	return json.Marshal(v)
}

func (v *TelevisionChannel) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TelevisionChannel"

	return json.Marshal(*v)
}
