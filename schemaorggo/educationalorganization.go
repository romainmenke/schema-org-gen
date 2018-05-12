package schemaorggo

import "encoding/json"

// EducationalOrganization see : https://schema.org/EducationalOrganization
type EducationalOrganization struct {
	Organization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	// types : Person
	Alumni *Person `json:"alumni,omitempty"`
}

func (v EducationalOrganization) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EducationalOrganization"

	return json.Marshal(v)
}

func (v *EducationalOrganization) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EducationalOrganization"

	return json.Marshal(*v)
}
