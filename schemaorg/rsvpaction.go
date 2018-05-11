package schemaorg

import "encoding/json"

// RsvpAction see : https://schema.org/RsvpAction
type RsvpAction struct {

typeContext

InformAction

// AdditionalNumberOfGuests see : /additionalNumberOfGuests
// If responding yes, the number of guests who will attend in addition to the invitee.
AdditionalNumberOfGuests float64 `json:"additionalNumberOfGuests"`

// Comment see : /comment
// Comments, typically from users.
Comment *Comment `json:"comment"`

// RsvpResponse see : /rsvpResponse
// The response (yes, no, maybe) to the RSVP.
RsvpResponse *RsvpResponseType `json:"rsvpResponse"`

}

func (v *RsvpAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RsvpAction"

	return json.Marshal(v)
}
