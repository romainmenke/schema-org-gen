package schemaorggo

import "encoding/json"

// PrependAction see : https://schema.org/PrependAction
type PrependAction struct {
	InsertAction

	typeContext

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation"`
}

func (v *PrependAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PrependAction"

	return json.Marshal(v)
}
