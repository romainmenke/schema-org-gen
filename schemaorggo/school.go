package schemaorggo

import "encoding/json"

// School see : https://schema.org/School
type School struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	Alumni *Person `json:"alumni,omitempty"` // types : Person

}

func (v School) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "School"

	return json.Marshal(v)
}

func (v *School) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "School"

	return json.Marshal(*v)
}
