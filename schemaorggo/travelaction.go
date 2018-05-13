package schemaorggo

import "encoding/json"

// TravelAction see : https://schema.org/TravelAction
type TravelAction struct {
	MoveAction

	typeContext

	// Distance see : https://schema.org/distance
	// The distance travelled, e.g. exercising or travelling.
	// types : Distance
	Distance []*Distance `json:"distance,omitempty"`
}

func (v TravelAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.MoveAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Distance
		if len(v.Distance) == 1 {
			value = v.Distance[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["distance"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TravelAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TravelAction"

	return data, nil
}

func (v TravelAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
