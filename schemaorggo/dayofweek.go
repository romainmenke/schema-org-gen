package schemaorggo

import "encoding/json"

// DayOfWeek see : https://schema.org/DayOfWeek
type DayOfWeek struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v DayOfWeek) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DayOfWeek"

	return json.Marshal(v)
}

func (v *DayOfWeek) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DayOfWeek"

	return json.Marshal(*v)
}
