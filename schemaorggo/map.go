package schemaorggo

import "encoding/json"

// Map see : https://schema.org/Map
type Map struct {
	CreativeWork

	typeContext

	// MapType see : https://schema.org/mapType
	// Indicates the kind of Map, from the MapCategoryType Enumeration.
	// types : MapCategoryType
	MapType []*MapCategoryType `json:"mapType,omitempty"`
}

func (v Map) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.MapType
		if len(v.MapType) == 1 {
			value = v.MapType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mapType"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Map) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Map"

	return data, nil
}

func (v Map) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
