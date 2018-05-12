package schemaorggo

import "encoding/json"

// ArriveAction see : https://schema.org/ArriveAction
type ArriveAction struct {
	MoveAction

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	// types : Place
	FromLocation *Place `json:"fromLocation,omitempty"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	// types : Place
	ToLocation *Place `json:"toLocation,omitempty"`
}

func (v ArriveAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ArriveAction"

	return json.Marshal(v)
}

func (v *ArriveAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ArriveAction"

	return json.Marshal(*v)
}
