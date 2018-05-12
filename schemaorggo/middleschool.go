package schemaorggo

import "encoding/json"

// MiddleSchool see : https://schema.org/MiddleSchool
type MiddleSchool struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	Alumni *Person `json:"alumni,omitempty"` // types : Person

}

func (v MiddleSchool) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MiddleSchool"

	return json.Marshal(v)
}

func (v *MiddleSchool) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MiddleSchool"

	return json.Marshal(*v)
}
