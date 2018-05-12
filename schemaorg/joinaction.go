package schemaorg

import "encoding/json"

// JoinAction see : https://schema.org/JoinAction
type JoinAction struct {

InteractAction

typeContext

// Event see : https://schema.org/event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event *Event `json:"event"`

}

func (v *JoinAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "JoinAction"

	return json.Marshal(v)
}
