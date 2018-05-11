package schemaorg

import "encoding/json"

// SportsOrganization see : https://schema.org/SportsOrganization
type SportsOrganization struct {

typeContext

Organization

// Sport see : https://schema.org/sport
// A type of sport (e.g. Baseball).
Sport interface{} `json:"sport"` // types : Text URL

}

func (v *SportsOrganization) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SportsOrganization"

	return json.Marshal(v)
}
