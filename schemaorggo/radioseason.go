package schemaorggo

import "encoding/json"

// RadioSeason see : https://schema.org/RadioSeason
type RadioSeason struct {
	CreativeWorkSeason

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate interface{} `json:"endDate,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	// types : Episode
	Episode *Episode `json:"episode,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes float64 `json:"numberOfEpisodes,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	// types : CreativeWorkSeries
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany *Organization `json:"productionCompany,omitempty"`

	// SeasonNumber see : https://schema.org/seasonNumber
	// Position of the season within an ordered group of seasons.
	// types : Integer Text
	SeasonNumber interface{} `json:"seasonNumber,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate interface{} `json:"startDate,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer *VideoObject `json:"trailer,omitempty"`
}

func (v RadioSeason) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioSeason"

	return json.Marshal(v)
}

func (v *RadioSeason) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RadioSeason"

	return json.Marshal(*v)
}
