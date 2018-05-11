package schemaorg

import "encoding/json"

// ConfirmAction see : https://schema.org/ConfirmAction
type ConfirmAction struct {

typeContext

InformAction

// Event see : https://schema.org/event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event *Event `json:"event"`

}

func (v *ConfirmAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ConfirmAction"

	return json.Marshal(v)
}
