package schemaorg

import "encoding/json"

// BroadcastEvent see : https://schema.org/BroadcastEvent
type BroadcastEvent struct {

typeContext

PublicationEvent

// BroadcastOfEvent see : /broadcastOfEvent
// The event being broadcast such as a sporting event or awards ceremony.
BroadcastOfEvent *Event `json:"broadcastOfEvent"`

// IsLiveBroadcast see : /isLiveBroadcast
// True is the broadcast is of a live event.
IsLiveBroadcast bool `json:"isLiveBroadcast"`

// VideoFormat see : /videoFormat
// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
VideoFormat string `json:"videoFormat"`

}

func (v *BroadcastEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BroadcastEvent"

	return json.Marshal(v)
}
