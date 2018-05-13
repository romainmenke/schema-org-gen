package schemaorggo

import "encoding/json"

// InviteAction see : https://schema.org/InviteAction
type InviteAction struct {
	CommunicateAction

	typeContext

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	// types : Event
	Event []*Event `json:"event,omitempty"`
}

func (v InviteAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CommunicateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Event
		if len(v.Event) == 1 {
			value = v.Event[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["event"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v InviteAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "InviteAction"

	return data, nil
}

func (v InviteAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
