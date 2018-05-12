package schemaorggo

import "encoding/json"

// TakeAction see : https://schema.org/TakeAction
type TakeAction struct {
	TransferAction

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	FromLocation *Place `json:"fromLocation,omitempty"` // types : Place

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation,omitempty"` // types : Place

}

func (v TakeAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TakeAction"

	return json.Marshal(v)
}

func (v *TakeAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TakeAction"

	return json.Marshal(*v)
}
