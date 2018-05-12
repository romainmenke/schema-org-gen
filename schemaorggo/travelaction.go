package schemaorggo

import "encoding/json"

// TravelAction see : https://schema.org/TravelAction
type TravelAction struct {
	MoveAction

	typeContext

	// Distance see : https://schema.org/distance
	// The distance travelled, e.g. exercising or travelling.
	// types : Distance
	Distance *Distance `json:"distance,omitempty"`
}

func (v TravelAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TravelAction"

	return json.Marshal(v)
}

func (v *TravelAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TravelAction"

	return json.Marshal(*v)
}
