package schemaorg

import "encoding/json"

// PrependAction see : https://schema.org/PrependAction
type PrependAction struct {

typeContext

InsertAction

// ToLocation see : /toLocation
// A sub property of location. The final location of the object or the agent after the action.
ToLocation *Place `json:"toLocation"`

}

func (v *PrependAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PrependAction"

	return json.Marshal(v)
}
