package schemaorggo

import "encoding/json"

// MapCategoryType see : https://schema.org/MapCategoryType
type MapCategoryType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v MapCategoryType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MapCategoryType"

	return json.Marshal(v)
}

func (v *MapCategoryType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MapCategoryType"

	return json.Marshal(*v)
}
