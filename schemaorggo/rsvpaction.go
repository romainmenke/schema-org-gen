package schemaorggo

import "encoding/json"

// RsvpAction see : https://schema.org/RsvpAction
type RsvpAction struct {
	InformAction

	typeContext

	// AdditionalNumberOfGuests see : https://schema.org/additionalNumberOfGuests
	// If responding yes, the number of guests who will attend in addition to the invitee.
	// types : Number
	AdditionalNumberOfGuests []float64 `json:"additionalNumberOfGuests,omitempty"`

	// Comment see : https://schema.org/comment
	// Comments, typically from users.
	// types : Comment
	Comment []*Comment `json:"comment,omitempty"`

	// RsvpResponse see : https://schema.org/rsvpResponse
	// The response (yes, no, maybe) to the RSVP.
	// types : RsvpResponseType
	RsvpResponse []*RsvpResponseType `json:"rsvpResponse,omitempty"`
}

func (v RsvpAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.InformAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AdditionalNumberOfGuests
		if len(v.AdditionalNumberOfGuests) == 1 {
			value = v.AdditionalNumberOfGuests[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalNumberOfGuests"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Comment
		if len(v.Comment) == 1 {
			value = v.Comment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["comment"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RsvpResponse
		if len(v.RsvpResponse) == 1 {
			value = v.RsvpResponse[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["rsvpResponse"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v RsvpAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "RsvpAction"

	return data, nil
}

func (v RsvpAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
