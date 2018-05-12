package schemaorggo

import "encoding/json"

// EducationalAudience see : https://schema.org/EducationalAudience
type EducationalAudience struct {
	Audience

	typeContext

	// EducationalRole see : https://schema.org/educationalRole
	// An educationalRole of an EducationalAudience.
	EducationalRole string `json:"educationalRole,omitempty"`
}

func (v EducationalAudience) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EducationalAudience"

	return json.Marshal(v)
}

func (v *EducationalAudience) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EducationalAudience"

	return json.Marshal(*v)
}
