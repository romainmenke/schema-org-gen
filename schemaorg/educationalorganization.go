package schemaorg

import "encoding/json"

// EducationalOrganization see : https://schema.org/EducationalOrganization
type EducationalOrganization struct {

typeContext

Organization

// Alumni see : /alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

}

func (v *EducationalOrganization) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EducationalOrganization"

	return json.Marshal(v)
}
