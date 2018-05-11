package schemaorg

import "encoding/json"

// BusinessEntityType see : https://schema.org/BusinessEntityType
type BusinessEntityType struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *BusinessEntityType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusinessEntityType"

	return json.Marshal(v)
}
