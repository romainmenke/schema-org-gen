package schemaorg

import "encoding/json"

// Audience see : https://schema.org/Audience
type Audience struct {

typeContext

Intangible

// AudienceType see : /audienceType
// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
AudienceType string `json:"audienceType"`

// GeographicArea see : /geographicArea
// The geographic area associated with the audience.
GeographicArea *AdministrativeArea `json:"geographicArea"`

}

func (v *Audience) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Audience"

	return json.Marshal(v)
}
