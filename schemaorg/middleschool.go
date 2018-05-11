package schemaorg

import "encoding/json"

// MiddleSchool see : https://schema.org/MiddleSchool
type MiddleSchool struct {

typeContext

EducationalOrganization

// Alumni see : https://schema.org/alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

}

func (v *MiddleSchool) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MiddleSchool"

	return json.Marshal(v)
}
