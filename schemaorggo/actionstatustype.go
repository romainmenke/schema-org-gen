package schemaorggo

import "encoding/json"

// ActionStatusType see : https://schema.org/ActionStatusType
type ActionStatusType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *ActionStatusType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ActionStatusType"

	return json.Marshal(v)
}
