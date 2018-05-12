package schemaorggo

import "encoding/json"

// BroadcastEvent see : https://schema.org/BroadcastEvent
type BroadcastEvent struct {
	PublicationEvent

	typeContext

	// BroadcastOfEvent see : https://schema.org/broadcastOfEvent
	// The event being broadcast such as a sporting event or awards ceremony.
	// types : Event
	BroadcastOfEvent *Event `json:"broadcastOfEvent,omitempty"`

	// IsLiveBroadcast see : https://schema.org/isLiveBroadcast
	// True is the broadcast is of a live event.
	// types : Boolean
	IsLiveBroadcast bool `json:"isLiveBroadcast,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat string `json:"videoFormat,omitempty"`
}

func (v BroadcastEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BroadcastEvent"

	return json.Marshal(v)
}

func (v *BroadcastEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BroadcastEvent"

	return json.Marshal(*v)
}
