package schemaorggo

import "encoding/json"

// TransferAction see : https://schema.org/TransferAction
type TransferAction struct {
	Action

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	FromLocation *Place `json:"fromLocation,omitempty"` // types : Place

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation,omitempty"` // types : Place

}

func (v TransferAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TransferAction"

	return json.Marshal(v)
}

func (v *TransferAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TransferAction"

	return json.Marshal(*v)
}
