package schemaorg

import "encoding/json"

// TVEpisode see : https://schema.org/TVEpisode
type TVEpisode struct {

typeContext

Episode

// CountryOfOrigin see : /countryOfOrigin
// The country of the principal offices of the production company or individual responsible for the movie or program.
CountryOfOrigin *Country `json:"countryOfOrigin"`

// SubtitleLanguage see : /subtitleLanguage
// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
SubtitleLanguage interface{} `json:"subtitleLanguage"` // types : Language Text

}

func (v *TVEpisode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TVEpisode"

	return json.Marshal(v)
}
