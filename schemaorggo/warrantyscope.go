package schemaorggo

import "encoding/json"

// WarrantyScope see : https://schema.org/WarrantyScope
type WarrantyScope struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v WarrantyScope) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WarrantyScope"

	return json.Marshal(v)
}

func (v *WarrantyScope) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WarrantyScope"

	return json.Marshal(*v)
}
