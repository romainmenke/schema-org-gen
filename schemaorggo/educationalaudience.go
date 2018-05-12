package schemaorggo

import "encoding/json"

// EducationalAudience see : https://schema.org/EducationalAudience
type EducationalAudience struct {
	Audience

	typeContext

	// EducationalRole see : https://schema.org/educationalRole
	// An educationalRole of an EducationalAudience.
	EducationalRole string `json:"educationalRole"`
}

func (v *EducationalAudience) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EducationalAudience"

	return json.Marshal(v)
}
