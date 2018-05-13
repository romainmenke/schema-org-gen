package schemaorggo

import "encoding/json"

// BroadcastService see : https://schema.org/BroadcastService
type BroadcastService struct {
	Service

	typeContext

	// BroadcastAffiliateOf see : https://schema.org/broadcastAffiliateOf
	// The media network(s) whose content is broadcast on this station.
	// types : Organization
	BroadcastAffiliateOf []*Organization `json:"broadcastAffiliateOf,omitempty"`

	// BroadcastDisplayName see : https://schema.org/broadcastDisplayName
	// The name displayed in the channel guide. For many US affiliates, it is the network name.
	// types : Text
	BroadcastDisplayName []string `json:"broadcastDisplayName,omitempty"`

	// BroadcastFrequency see : http://pending.schema.org/broadcastFrequency
	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	// types : BroadcastFrequencySpecification Text
	BroadcastFrequency []interface{} `json:"broadcastFrequency,omitempty"`

	// BroadcastTimezone see : https://schema.org/broadcastTimezone
	// The timezone in ISO 8601 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601) for which the service bases its broadcasts
	// types : Text
	BroadcastTimezone []string `json:"broadcastTimezone,omitempty"`

	// Broadcaster see : https://schema.org/broadcaster
	// The organization owning or operating the broadcast service.
	// types : Organization
	Broadcaster []*Organization `json:"broadcaster,omitempty"`

	// HasBroadcastChannel see : http://pending.schema.org/hasBroadcastChannel
	// A broadcast channel of a broadcast service. Inverse property: providesBroadcastService (see: https://schema.org/providesBroadcastService).
	// types : BroadcastChannel
	HasBroadcastChannel []*BroadcastChannel `json:"hasBroadcastChannel,omitempty"`

	// ParentService see : https://schema.org/parentService
	// A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	// types : BroadcastService
	ParentService []*BroadcastService `json:"parentService,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat []string `json:"videoFormat,omitempty"`
}

func (v BroadcastService) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Service.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BroadcastAffiliateOf
		if len(v.BroadcastAffiliateOf) == 1 {
			value = v.BroadcastAffiliateOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastAffiliateOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BroadcastDisplayName
		if len(v.BroadcastDisplayName) == 1 {
			value = v.BroadcastDisplayName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastDisplayName"] = json.RawMessage(b)
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
		var value interface{} = v.BroadcastTimezone
		if len(v.BroadcastTimezone) == 1 {
			value = v.BroadcastTimezone[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcastTimezone"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Broadcaster
		if len(v.Broadcaster) == 1 {
			value = v.Broadcaster[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broadcaster"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasBroadcastChannel
		if len(v.HasBroadcastChannel) == 1 {
			value = v.HasBroadcastChannel[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasBroadcastChannel"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ParentService
		if len(v.ParentService) == 1 {
			value = v.ParentService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["parentService"] = json.RawMessage(b)
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

func (v BroadcastService) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BroadcastService"

	return data, nil
}

func (v BroadcastService) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
