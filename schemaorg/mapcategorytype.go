package schemaorg

import "encoding/json"

// MapCategoryType see : https://schema.org/MapCategoryType
type MapCategoryType struct {

typeContext

Enumeration

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *MapCategoryType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MapCategoryType"

	return json.Marshal(v)
}
