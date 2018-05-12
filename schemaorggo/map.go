package schemaorggo

import "encoding/json"

// Map see : https://schema.org/Map
type Map struct {
	CreativeWork

	typeContext

	// MapType see : https://schema.org/mapType
	// Indicates the kind of Map, from the MapCategoryType Enumeration.
	MapType *MapCategoryType `json:"mapType,omitempty"`
}

func (v Map) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Map"

	return json.Marshal(v)
}

func (v *Map) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Map"

	return json.Marshal(*v)
}
