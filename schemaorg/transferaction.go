package schemaorg

import "encoding/json"

// TransferAction see : https://schema.org/TransferAction
type TransferAction struct {

typeContext

Action

// FromLocation see : /fromLocation
// A sub property of location. The original location of the object or the agent before the action.
FromLocation *Place `json:"fromLocation"`

// ToLocation see : /toLocation
// A sub property of location. The final location of the object or the agent after the action.
ToLocation *Place `json:"toLocation"`

}

func (v *TransferAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TransferAction"

	return json.Marshal(v)
}
