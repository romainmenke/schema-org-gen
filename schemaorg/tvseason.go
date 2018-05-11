package schemaorg

import "encoding/json"

// TVSeason see : https://schema.org/TVSeason
type TVSeason struct {

typeContext

CreativeWork

// CountryOfOrigin see : /countryOfOrigin
// The country of the principal offices of the production company or individual responsible for the movie or program.
CountryOfOrigin *Country `json:"countryOfOrigin"`

}

func (v *TVSeason) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TVSeason"

	return json.Marshal(v)
}
