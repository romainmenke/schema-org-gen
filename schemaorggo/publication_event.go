package schemaorggo

import "encoding/json"

// PublicationEvent see : https://schema.org/PublicationEvent
type PublicationEvent struct {
	Event

	typeContext

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	// types : Boolean
	IsAccessibleForFree []bool `json:"isAccessibleForFree,omitempty"`

	// PublishedBy see : https://schema.org/publishedBy
	// An agent associated with the publication event.
	// types : Organization Person
	PublishedBy []interface{} `json:"publishedBy,omitempty"`

	// PublishedOn see : https://schema.org/publishedOn
	// A broadcast service associated with the publication event.
	// types : BroadcastService
	PublishedOn []*BroadcastService `json:"publishedOn,omitempty"`
}

func (v PublicationEvent) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Event.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.IsAccessibleForFree
		if len(v.IsAccessibleForFree) == 1 {
			value = v.IsAccessibleForFree[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isAccessibleForFree"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PublishedBy
		if len(v.PublishedBy) == 1 {
			value = v.PublishedBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publishedBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PublishedOn
		if len(v.PublishedOn) == 1 {
			value = v.PublishedOn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publishedOn"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PublicationEvent) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PublicationEvent"

	return data, nil
}

func (v PublicationEvent) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
