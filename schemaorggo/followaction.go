package schemaorggo

import "encoding/json"

// FollowAction see : https://schema.org/FollowAction
type FollowAction struct {
	InteractAction

	typeContext

	// Followee see : https://schema.org/followee
	// A sub property of object. The person or organization being followed.
	// types : Organization Person
	Followee []interface{} `json:"followee,omitempty"`
}

func (v FollowAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.InteractAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Followee
		if len(v.Followee) == 1 {
			value = v.Followee[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["followee"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v FollowAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "FollowAction"

	return data, nil
}

func (v FollowAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
