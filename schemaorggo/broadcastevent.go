package schemaorggo

import "encoding/json"

// BroadcastEvent see : https://schema.org/BroadcastEvent
type BroadcastEvent struct {
	PublicationEvent

	typeContext

	// BroadcastOfEvent see : https://schema.org/broadcastOfEvent
	// The event being broadcast such as a sporting event or awards ceremony.
	// types : Event
	BroadcastOfEvent []*Event `json:"broadcastOfEvent,omitempty"`

	// IsLiveBroadcast see : https://schema.org/isLiveBroadcast
	// True is the broadcast is of a live event.
	// types : Boolean
	IsLiveBroadcast []bool `json:"isLiveBroadcast,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat []string `json:"videoFormat,omitempty"`
}

func (v BroadcastEvent) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PublicationEvent.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BroadcastOfEvent
		if len(v.BroadcastOfEvent) == 1 {
			value = v.BroadcastOfEvent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastOfEvent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsLiveBroadcast
		if len(v.IsLiveBroadcast) == 1 {
			value = v.IsLiveBroadcast[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isLiveBroadcast"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VideoFormat
		if len(v.VideoFormat) == 1 {
			value = v.VideoFormat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["videoFormat"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BroadcastEvent) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BroadcastEvent"

	return data, nil
}

func (v BroadcastEvent) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
