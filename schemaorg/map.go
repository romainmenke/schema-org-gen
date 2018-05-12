package schemaorg

import "encoding/json"

// Map see : https://schema.org/Map
type Map struct {

CreativeWork

typeContext

// MapType see : https://schema.org/mapType
// Indicates the kind of Map, from the MapCategoryType Enumeration.
MapType *MapCategoryType `json:"mapType"`

}

func (v *Map) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Map"

	return json.Marshal(v)
}
