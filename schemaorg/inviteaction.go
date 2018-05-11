package schemaorg

import "encoding/json"

// InviteAction see : https://schema.org/InviteAction
type InviteAction struct {

typeContext

CommunicateAction

// Event see : /event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event *Event `json:"event"`

}

func (v *InviteAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InviteAction"

	return json.Marshal(v)
}
