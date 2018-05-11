package schemaorg

import "encoding/json"

// CreativeWorkSeason see : https://schema.org/CreativeWorkSeason
type CreativeWorkSeason struct {

typeContext

CreativeWork

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// EndDate see : /endDate
// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
EndDate interface{} `json:"endDate"` // types : Date DateTime

// Episode see : /episode
// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
Episode *Episode `json:"episode"`

// NumberOfEpisodes see : /numberOfEpisodes
// The number of episodes in this season or series.
NumberOfEpisodes int `json:"numberOfEpisodes"`

// PartOfSeries see : /partOfSeries
// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
PartOfSeries *CreativeWorkSeries `json:"partOfSeries"`

// ProductionCompany see : /productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc.
ProductionCompany *Organization `json:"productionCompany"`

// SeasonNumber see : /seasonNumber
// Position of the season within an ordered group of seasons.
SeasonNumber interface{} `json:"seasonNumber"` // types : Integer Text

// StartDate see : /startDate
// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
StartDate interface{} `json:"startDate"` // types : Date DateTime

// Trailer see : /trailer
// The trailer of a movie or tv/radio series, season, episode, etc.
Trailer *VideoObject `json:"trailer"`

}

func (v *CreativeWorkSeason) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CreativeWorkSeason"

	return json.Marshal(v)
}
