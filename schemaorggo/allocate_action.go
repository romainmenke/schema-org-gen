package schemaorggo

import "encoding/json"

// AllocateAction see : https://schema.org/AllocateAction
type AllocateAction struct {
	OrganizeAction

	typeContext

	// Purpose see : https://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	// types : MedicalDevicePurpose Thing
	Purpose []interface{} `json:"purpose,omitempty"`
}

func (v AllocateAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.OrganizeAction.intoMap(intop)

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

	*intop = into

	return nil
}

func (v AllocateAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AllocateAction"

	return data, nil
}

func (v AllocateAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
