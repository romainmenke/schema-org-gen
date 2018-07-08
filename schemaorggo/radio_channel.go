package schemaorggo

import "encoding/json"

// RadioChannel see : https://schema.org/RadioChannel
type RadioChannel struct {
	BroadcastChannel

	typeContext

	// BroadcastChannelId see : https://schema.org/broadcastChannelId
	// The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
	// types : Text
	BroadcastChannelId []string `json:"broadcastChannelId,omitempty"`

	// BroadcastFrequency see : https://pending.schema.org/broadcastFrequency
	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	// types : BroadcastFrequencySpecification Text
	BroadcastFrequency []interface{} `json:"broadcastFrequency,omitempty"`

	// BroadcastServiceTier see : https://schema.org/broadcastServiceTier
	// The type of service required to have access to the channel (e.g. Standard or Premium).
	// types : Text
	BroadcastServiceTier []string `json:"broadcastServiceTier,omitempty"`

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	// types : Text URL
	Genre []string `json:"genre,omitempty"`

	// InBroadcastLineup see : https://schema.org/inBroadcastLineup
	// The CableOrSatelliteService offering the channel.
	// types : CableOrSatelliteService
	InBroadcastLineup []*CableOrSatelliteService `json:"inBroadcastLineup,omitempty"`

	// ProvidesBroadcastService see : https://schema.org/providesBroadcastService
	// The BroadcastService offered on this channel. Inverse property: hasBroadcastChannel (see: https://schema.orghttps://pending.schema.org/hasBroadcastChannel).
	// types : BroadcastService
	ProvidesBroadcastService []*BroadcastService `json:"providesBroadcastService,omitempty"`
}

func (v RadioChannel) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.BroadcastChannel.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BroadcastChannelId
		if len(v.BroadcastChannelId) == 1 {
			value = v.BroadcastChannelId[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastChannelId"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BroadcastFrequency
		if len(v.BroadcastFrequency) == 1 {
			value = v.BroadcastFrequency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastFrequency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BroadcastServiceTier
		if len(v.BroadcastServiceTier) == 1 {
			value = v.BroadcastServiceTier[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastServiceTier"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Genre
		if len(v.Genre) == 1 {
			value = v.Genre[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["genre"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InBroadcastLineup
		if len(v.InBroadcastLineup) == 1 {
			value = v.InBroadcastLineup[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inBroadcastLineup"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProvidesBroadcastService
		if len(v.ProvidesBroadcastService) == 1 {
			value = v.ProvidesBroadcastService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["providesBroadcastService"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v RadioChannel) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "RadioChannel"

	return data, nil
}

func (v RadioChannel) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
