package schemaorggo

import "encoding/json"

// LeaveAction see : https://schema.org/LeaveAction
type LeaveAction struct {
	InteractAction

	typeContext

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	Event *Event `json:"event,omitempty"`
}

func (v LeaveAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LeaveAction"

	return json.Marshal(v)
}

func (v *LeaveAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LeaveAction"

	return json.Marshal(*v)
}
