package schemaorg

import "encoding/json"

// Preschool see : https://schema.org/Preschool
type Preschool struct {

typeContext

EducationalOrganization

// Alumni see : https://schema.org/alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

}

func (v *Preschool) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Preschool"

	return json.Marshal(v)
}
