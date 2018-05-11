package schemaorg

import "encoding/json"

// DayOfWeek see : https://schema.org/DayOfWeek
type DayOfWeek struct {

typeContext

Enumeration

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *DayOfWeek) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DayOfWeek"

	return json.Marshal(v)
}
