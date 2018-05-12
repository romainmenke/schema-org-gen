package schemaorggo

import "encoding/json"

// JoinAction see : https://schema.org/JoinAction
type JoinAction struct {
	InteractAction

	typeContext

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	Event *Event `json:"event,omitempty"`
}

func (v JoinAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "JoinAction"

	return json.Marshal(v)
}

func (v *JoinAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "JoinAction"

	return json.Marshal(*v)
}
