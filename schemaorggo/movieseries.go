package schemaorggo

import "encoding/json"

// MovieSeries see : https://schema.org/MovieSeries
type MovieSeries struct {
	CreativeWorkSeries

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy,omitempty"` // types : MusicGroup Person

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	ProductionCompany *Organization `json:"productionCompany,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer *VideoObject `json:"trailer,omitempty"`
}

func (v MovieSeries) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MovieSeries"

	return json.Marshal(v)
}

func (v *MovieSeries) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MovieSeries"

	return json.Marshal(*v)
}
