package schemaorg

import "encoding/json"

// ElementarySchool see : https://schema.org/ElementarySchool
type ElementarySchool struct {

typeContext

EducationalOrganization

// Alumni see : https://schema.org/alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

}

func (v *ElementarySchool) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ElementarySchool"

	return json.Marshal(v)
}
