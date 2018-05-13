package schemaorggo

import "encoding/json"

// SearchAction see : https://schema.org/SearchAction
type SearchAction struct {
	Action

	typeContext

	// Query see : https://schema.org/query
	// A sub property of instrument. The query used on this action.
	// types : Text
	Query []string `json:"query,omitempty"`
}

func (v SearchAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Action.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Query
		if len(v.Query) == 1 {
			value = v.Query[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["query"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SearchAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SearchAction"

	return data, nil
}

func (v SearchAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
