package schemaorggo

import "encoding/json"

// RadioEpisode see : https://schema.org/RadioEpisode
type RadioEpisode struct {
	Episode

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// EpisodeNumber see : https://schema.org/episodeNumber
	// Position of the episode within an ordered group of episodes.
	// types : Integer Text
	EpisodeNumber interface{} `json:"episodeNumber,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy interface{} `json:"musicBy,omitempty"`

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	// types : CreativeWorkSeason
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	// types : CreativeWorkSeries
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany *Organization `json:"productionCompany,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer *VideoObject `json:"trailer,omitempty"`
}

func (v RadioEpisode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioEpisode"

	return json.Marshal(v)
}

func (v *RadioEpisode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RadioEpisode"

	return json.Marshal(*v)
}
