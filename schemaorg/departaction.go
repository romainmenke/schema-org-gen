package schemaorg

import "encoding/json"

// DepartAction see : https://schema.org/DepartAction
type DepartAction struct {

typeContext

MoveAction

// FromLocation see : /fromLocation
// A sub property of location. The original location of the object or the agent before the action.
FromLocation *Place `json:"fromLocation"`

// ToLocation see : /toLocation
// A sub property of location. The final location of the object or the agent after the action.
ToLocation *Place `json:"toLocation"`

}

func (v *DepartAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DepartAction"

	return json.Marshal(v)
}
