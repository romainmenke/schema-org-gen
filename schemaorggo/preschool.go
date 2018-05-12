package schemaorggo

import "encoding/json"

// Preschool see : https://schema.org/Preschool
type Preschool struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	Alumni *Person `json:"alumni,omitempty"`
}

func (v Preschool) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Preschool"

	return json.Marshal(v)
}

func (v *Preschool) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Preschool"

	return json.Marshal(*v)
}
