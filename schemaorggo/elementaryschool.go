package schemaorggo

import "encoding/json"

// ElementarySchool see : https://schema.org/ElementarySchool
type ElementarySchool struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	// types : Person
	Alumni *Person `json:"alumni,omitempty"`
}

func (v ElementarySchool) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ElementarySchool"

	return json.Marshal(v)
}

func (v *ElementarySchool) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ElementarySchool"

	return json.Marshal(*v)
}
