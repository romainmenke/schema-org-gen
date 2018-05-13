package schemaorggo

import "encoding/json"

// PerformAction see : https://schema.org/PerformAction
type PerformAction struct {
	PlayAction

	typeContext

	// EntertainmentBusiness see : https://schema.org/entertainmentBusiness
	// A sub property of location. The entertainment business where the action occurred.
	// types : EntertainmentBusiness
	EntertainmentBusiness []*EntertainmentBusiness `json:"entertainmentBusiness,omitempty"`
}

func (v PerformAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PlayAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.EntertainmentBusiness
		if len(v.EntertainmentBusiness) == 1 {
			value = v.EntertainmentBusiness[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["entertainmentBusiness"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PerformAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PerformAction"

	return data, nil
}

func (v PerformAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
