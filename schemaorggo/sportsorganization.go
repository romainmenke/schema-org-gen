package schemaorggo

import "encoding/json"

// SportsOrganization see : https://schema.org/SportsOrganization
type SportsOrganization struct {
	Organization

	typeContext

	// Sport see : https://schema.org/sport
	// A type of sport (e.g. Baseball).
	// types : Text URL
	Sport string `json:"sport,omitempty"`
}

func (v SportsOrganization) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SportsOrganization"

	return json.Marshal(v)
}

func (v *SportsOrganization) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SportsOrganization"

	return json.Marshal(*v)
}
