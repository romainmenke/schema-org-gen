package schemaorg

import "encoding/json"

// WarrantyScope see : https://schema.org/WarrantyScope
type WarrantyScope struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *WarrantyScope) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WarrantyScope"

	return json.Marshal(v)
}
