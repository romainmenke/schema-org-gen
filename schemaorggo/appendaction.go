package schemaorggo

import "encoding/json"

// AppendAction see : https://schema.org/AppendAction
type AppendAction struct {
	InsertAction

	typeContext

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation"`
}

func (v *AppendAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AppendAction"

	return json.Marshal(v)
}
