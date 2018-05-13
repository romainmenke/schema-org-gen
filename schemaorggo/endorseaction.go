package schemaorggo

import "encoding/json"

// EndorseAction see : https://schema.org/EndorseAction
type EndorseAction struct {
	ReactAction

	typeContext

	// Endorsee see : https://schema.org/endorsee
	// A sub property of participant. The person/organization being supported.
	// types : Organization Person
	Endorsee []interface{} `json:"endorsee,omitempty"`
}

func (v EndorseAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ReactAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Endorsee
		if len(v.Endorsee) == 1 {
			value = v.Endorsee[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endorsee"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v EndorseAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "EndorseAction"

	return data, nil
}

func (v EndorseAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
