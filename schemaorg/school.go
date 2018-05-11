package schemaorg

import "encoding/json"

// School see : https://schema.org/School
type School struct {

typeContext

EducationalOrganization

// Alumni see : https://schema.org/alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

}

func (v *School) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "School"

	return json.Marshal(v)
}
