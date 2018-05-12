package schemaorg

import "encoding/json"

// PlayAction see : https://schema.org/PlayAction
type PlayAction struct {

Action

typeContext

// Audience see : https://schema.org/audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience *Audience `json:"audience"`

// Event see : https://schema.org/event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event *Event `json:"event"`

}

func (v *PlayAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PlayAction"

	return json.Marshal(v)
}
