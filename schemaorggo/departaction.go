package schemaorggo

import "encoding/json"

// DepartAction see : https://schema.org/DepartAction
type DepartAction struct {
	MoveAction

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	FromLocation *Place `json:"fromLocation,omitempty"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation,omitempty"`
}

func (v DepartAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DepartAction"

	return json.Marshal(v)
}

func (v *DepartAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DepartAction"

	return json.Marshal(*v)
}
