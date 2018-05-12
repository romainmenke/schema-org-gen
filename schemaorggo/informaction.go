package schemaorggo

import "encoding/json"

// InformAction see : https://schema.org/InformAction
type InformAction struct {
	CommunicateAction

	typeContext

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	Event *Event `json:"event"`
}

func (v *InformAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InformAction"

	return json.Marshal(v)
}
