package schemaorggo

import "encoding/json"

// ConfirmAction see : https://schema.org/ConfirmAction
type ConfirmAction struct {
	InformAction

	typeContext

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	Event *Event `json:"event,omitempty"` // types : Event

}

func (v ConfirmAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ConfirmAction"

	return json.Marshal(v)
}

func (v *ConfirmAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ConfirmAction"

	return json.Marshal(*v)
}
