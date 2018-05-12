package schemaorggo

import "encoding/json"

// RadioSeries see : https://schema.org/RadioSeries
type RadioSeries struct {
	CreativeWorkSeries

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// ContainsSeason see : https://schema.org/containsSeason
	// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
	// types : CreativeWorkSeason
	ContainsSeason *CreativeWorkSeason `json:"containsSeason,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	// types : Episode
	Episode *Episode `json:"episode,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy interface{} `json:"musicBy,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes float64 `json:"numberOfEpisodes,omitempty"`

	// NumberOfSeasons see : https://schema.org/numberOfSeasons
	// The number of seasons in this series.
	// types : Integer
	NumberOfSeasons float64 `json:"numberOfSeasons,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany *Organization `json:"productionCompany,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer *VideoObject `json:"trailer,omitempty"`
}

func (v RadioSeries) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioSeries"

	return json.Marshal(v)
}

func (v *RadioSeries) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RadioSeries"

	return json.Marshal(*v)
}
