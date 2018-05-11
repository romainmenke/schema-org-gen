package schemaorg

import "encoding/json"

// MoveAction see : https://schema.org/MoveAction
type MoveAction struct {

typeContext

Action

// FromLocation see : /fromLocation
// A sub property of location. The original location of the object or the agent before the action.
FromLocation *Place `json:"fromLocation"`

// ToLocation see : /toLocation
// A sub property of location. The final location of the object or the agent after the action.
ToLocation *Place `json:"toLocation"`

}

func (v *MoveAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MoveAction"

	return json.Marshal(v)
}
