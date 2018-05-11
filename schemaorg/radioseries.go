package schemaorg

import "encoding/json"

// RadioSeries see : https://schema.org/RadioSeries
type RadioSeries struct {

typeContext

CreativeWorkSeries

// Actor see : https://schema.org/actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// ContainsSeason see : https://schema.org/containsSeason
// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
ContainsSeason *CreativeWorkSeason `json:"containsSeason"`

// Director see : https://schema.org/director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// Episode see : https://schema.org/episode
// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
Episode *Episode `json:"episode"`

// MusicBy see : https://schema.org/musicBy
// The composer of the soundtrack.
MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
// The number of episodes in this season or series.
NumberOfEpisodes int `json:"numberOfEpisodes"`

// NumberOfSeasons see : https://schema.org/numberOfSeasons
// The number of seasons in this series.
NumberOfSeasons int `json:"numberOfSeasons"`

// ProductionCompany see : https://schema.org/productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc.
ProductionCompany *Organization `json:"productionCompany"`

// Trailer see : https://schema.org/trailer
// The trailer of a movie or tv/radio series, season, episode, etc.
Trailer *VideoObject `json:"trailer"`

}

func (v *RadioSeries) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioSeries"

	return json.Marshal(v)
}
