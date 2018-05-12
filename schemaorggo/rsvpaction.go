package schemaorggo

import "encoding/json"

// RsvpAction see : https://schema.org/RsvpAction
type RsvpAction struct {
	InformAction

	typeContext

	// AdditionalNumberOfGuests see : https://schema.org/additionalNumberOfGuests
	// If responding yes, the number of guests who will attend in addition to the invitee.
	AdditionalNumberOfGuests float64 `json:"additionalNumberOfGuests,omitempty"`

	// Comment see : https://schema.org/comment
	// Comments, typically from users.
	Comment *Comment `json:"comment,omitempty"`

	// RsvpResponse see : https://schema.org/rsvpResponse
	// The response (yes, no, maybe) to the RSVP.
	RsvpResponse *RsvpResponseType `json:"rsvpResponse,omitempty"`
}

func (v RsvpAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RsvpAction"

	return json.Marshal(v)
}

func (v *RsvpAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RsvpAction"

	return json.Marshal(*v)
}
