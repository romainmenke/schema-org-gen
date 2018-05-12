package schemaorggo

import "encoding/json"

// ItemAvailability see : https://schema.org/ItemAvailability
type ItemAvailability struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *ItemAvailability) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ItemAvailability"

	return json.Marshal(v)
}
