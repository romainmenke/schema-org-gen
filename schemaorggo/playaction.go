package schemaorggo

import "encoding/json"

// PlayAction see : https://schema.org/PlayAction
type PlayAction struct {
	Action

	typeContext

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience *Audience `json:"audience,omitempty"`

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	// types : Event
	Event *Event `json:"event,omitempty"`
}

func (v PlayAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PlayAction"

	return json.Marshal(v)
}

func (v *PlayAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PlayAction"

	return json.Marshal(*v)
}
