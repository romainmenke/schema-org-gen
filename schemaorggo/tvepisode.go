package schemaorggo

import "encoding/json"

// TVEpisode see : https://schema.org/TVEpisode
type TVEpisode struct {
	Episode

	typeContext

	// CountryOfOrigin see : https://schema.org/countryOfOrigin
	// The country of the principal offices of the production company or individual responsible for the movie or program.
	CountryOfOrigin *Country `json:"countryOfOrigin,omitempty"` // types : Country

	// SubtitleLanguage see : https://schema.org/subtitleLanguage
	// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty"` // types : Language Text

}

func (v TVEpisode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TVEpisode"

	return json.Marshal(v)
}

func (v *TVEpisode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TVEpisode"

	return json.Marshal(*v)
}
