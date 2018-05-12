package schemaorggo

import "encoding/json"

// CollegeOrUniversity see : https://schema.org/CollegeOrUniversity
type CollegeOrUniversity struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	Alumni *Person `json:"alumni,omitempty"`
}

func (v CollegeOrUniversity) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CollegeOrUniversity"

	return json.Marshal(v)
}

func (v *CollegeOrUniversity) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CollegeOrUniversity"

	return json.Marshal(*v)
}
