package schemaorggo

import "encoding/json"

// MoveAction see : https://schema.org/MoveAction
type MoveAction struct {
	Action

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	FromLocation *Place `json:"fromLocation"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation"`
}

func (v *MoveAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MoveAction"

	return json.Marshal(v)
}
