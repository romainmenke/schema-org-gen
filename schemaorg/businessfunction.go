package schemaorg

import "encoding/json"

// BusinessFunction see : https://schema.org/BusinessFunction
type BusinessFunction struct {

Enumeration

typeContext

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *BusinessFunction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusinessFunction"

	return json.Marshal(v)
}
