package schemaorggo

import "encoding/json"

// Episode see : https://schema.org/Episode
type Episode struct {
	CreativeWork

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor,omitempty"` // types : Person

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director,omitempty"` // types : Person

	// EpisodeNumber see : https://schema.org/episodeNumber
	// Position of the episode within an ordered group of episodes.
	EpisodeNumber interface{} `json:"episodeNumber,omitempty"` // types : Integer Text

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy,omitempty"` // types : MusicGroup Person

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason,omitempty"` // types : CreativeWorkSeason

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty"` // types : CreativeWorkSeries

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	ProductionCompany *Organization `json:"productionCompany,omitempty"` // types : Organization

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer *VideoObject `json:"trailer,omitempty"` // types : VideoObject

}

func (v Episode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Episode"

	return json.Marshal(v)
}

func (v *Episode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Episode"

	return json.Marshal(*v)
}
