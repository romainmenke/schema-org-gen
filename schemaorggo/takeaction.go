package schemaorggo

import "encoding/json"

// TakeAction see : https://schema.org/TakeAction
type TakeAction struct {
	TransferAction

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	FromLocation *Place `json:"fromLocation"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation"`
}

func (v *TakeAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TakeAction"

	return json.Marshal(v)
}
