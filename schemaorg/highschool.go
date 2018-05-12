package schemaorg

import "encoding/json"

// HighSchool see : https://schema.org/HighSchool
type HighSchool struct {

EducationalOrganization

typeContext

// Alumni see : https://schema.org/alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

}

func (v *HighSchool) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HighSchool"

	return json.Marshal(v)
}
