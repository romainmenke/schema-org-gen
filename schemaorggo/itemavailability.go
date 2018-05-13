package schemaorggo

import "encoding/json"

// ItemAvailability see : https://schema.org/ItemAvailability
type ItemAvailability struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy []interface{} `json:"supersededBy,omitempty"`
}

func (v ItemAvailability) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Enumeration.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.SupersededBy
		if len(v.SupersededBy) == 1 {
			value = v.SupersededBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["supersededBy"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ItemAvailability) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ItemAvailability"

	return data, nil
}

func (v ItemAvailability) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
