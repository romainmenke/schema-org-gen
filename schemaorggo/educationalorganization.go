package schemaorggo

import "encoding/json"

// EducationalOrganization see : https://schema.org/EducationalOrganization
type EducationalOrganization struct {
	Organization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	Alumni *Person `json:"alumni"`
}

func (v *EducationalOrganization) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EducationalOrganization"

	return json.Marshal(v)
}
