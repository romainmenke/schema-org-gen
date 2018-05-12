package schemaorggo

import "encoding/json"

// HighSchool see : https://schema.org/HighSchool
type HighSchool struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	// types : Person
	Alumni *Person `json:"alumni,omitempty"`
}

func (v HighSchool) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HighSchool"

	return json.Marshal(v)
}

func (v *HighSchool) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HighSchool"

	return json.Marshal(*v)
}
