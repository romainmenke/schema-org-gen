package schemaorggo

import "encoding/json"

// Audience see : https://schema.org/Audience
type Audience struct {
	Intangible

	typeContext

	// AudienceType see : https://schema.org/audienceType
	// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	AudienceType string `json:"audienceType,omitempty"`

	// GeographicArea see : https://schema.org/geographicArea
	// The geographic area associated with the audience.
	GeographicArea *AdministrativeArea `json:"geographicArea,omitempty"`
}

func (v Audience) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Audience"

	return json.Marshal(v)
}

func (v *Audience) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Audience"

	return json.Marshal(*v)
}
