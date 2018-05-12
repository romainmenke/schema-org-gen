package schemaorggo

import "encoding/json"

// Movie see : https://schema.org/Movie
type Movie struct {
	CreativeWork

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor"`

	// CountryOfOrigin see : https://schema.org/countryOfOrigin
	// The country of the principal offices of the production company or individual responsible for the movie or program.
	CountryOfOrigin *Country `json:"countryOfOrigin"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	Duration *Duration `json:"duration"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	ProductionCompany *Organization `json:"productionCompany"`

	// SubtitleLanguage see : https://schema.org/subtitleLanguage
	// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	SubtitleLanguage interface{} `json:"subtitleLanguage"` // types : Language Text

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer *VideoObject `json:"trailer"`
}

func (v *Movie) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Movie"

	return json.Marshal(v)
}
