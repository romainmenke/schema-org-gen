package schemaorggo

import "encoding/json"

// EventStatusType see : https://schema.org/EventStatusType
type EventStatusType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v EventStatusType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EventStatusType"

	return json.Marshal(v)
}

func (v *EventStatusType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EventStatusType"

	return json.Marshal(*v)
}
