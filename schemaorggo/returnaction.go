package schemaorggo

import "encoding/json"

// ReturnAction see : https://schema.org/ReturnAction
type ReturnAction struct {
	TransferAction

	typeContext

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient []interface{} `json:"recipient,omitempty"`
}

func (v ReturnAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TransferAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Recipient
		if len(v.Recipient) == 1 {
			value = v.Recipient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipient"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReturnAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReturnAction"

	return data, nil
}

func (v ReturnAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
