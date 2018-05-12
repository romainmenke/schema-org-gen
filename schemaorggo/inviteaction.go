package schemaorggo

import "encoding/json"

// InviteAction see : https://schema.org/InviteAction
type InviteAction struct {
	CommunicateAction

	typeContext

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	Event *Event `json:"event,omitempty"` // types : Event

}

func (v InviteAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InviteAction"

	return json.Marshal(v)
}

func (v *InviteAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "InviteAction"

	return json.Marshal(*v)
}
