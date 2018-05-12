package schemaorggo

import "encoding/json"

// InsertAction see : https://schema.org/InsertAction
type InsertAction struct {
	AddAction

	typeContext

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation"`
}

func (v *InsertAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InsertAction"

	return json.Marshal(v)
}
