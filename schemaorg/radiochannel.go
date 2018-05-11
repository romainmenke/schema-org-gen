package schemaorg

import "encoding/json"

// RadioChannel see : https://schema.org/RadioChannel
type RadioChannel struct {

typeContext

BroadcastChannel

// BroadcastChannelId see : /broadcastChannelId
// The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
BroadcastChannelId string `json:"broadcastChannelId"`

// BroadcastFrequency see : http://pending.schema.org/broadcastFrequency
// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. "87 FM".
BroadcastFrequency interface{} `json:"broadcastFrequency"` // types : BroadcastFrequencySpecification Text

// BroadcastServiceTier see : /broadcastServiceTier
// The type of service required to have access to the channel (e.g. Standard or Premium).
BroadcastServiceTier string `json:"broadcastServiceTier"`

// Genre see : /genre
// Genre of the creative work, broadcast channel or group.
Genre interface{} `json:"genre"` // types : Text URL

// InBroadcastLineup see : /inBroadcastLineup
// The CableOrSatelliteService offering the channel.
InBroadcastLineup *CableOrSatelliteService `json:"inBroadcastLineup"`

// ProvidesBroadcastService see : /providesBroadcastService
// The BroadcastService offered on this channel. Inverse property: hasBroadcastChannel (see: https://schema.orghttp://pending.schema.org/hasBroadcastChannel).
ProvidesBroadcastService *BroadcastService `json:"providesBroadcastService"`

}

func (v *RadioChannel) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioChannel"

	return json.Marshal(v)
}
