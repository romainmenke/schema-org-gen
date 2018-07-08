package schemaorggo

import "encoding/json"

// PayAction see : https://schema.org/PayAction
type PayAction struct {
	TradeAction

	typeContext

	// Purpose see : https://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	// types : MedicalDevicePurpose Thing
	Purpose []interface{} `json:"purpose,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient []interface{} `json:"recipient,omitempty"`
}

func (v PayAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TradeAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Purpose
		if len(v.Purpose) == 1 {
			value = v.Purpose[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["purpose"] = json.RawMessage(b)
		}
	}

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

func (v PayAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PayAction"

	return data, nil
}

func (v PayAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
