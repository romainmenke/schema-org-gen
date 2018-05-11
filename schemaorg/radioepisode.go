package schemaorg

import "encoding/json"

// RadioEpisode see : https://schema.org/RadioEpisode
type RadioEpisode struct {

typeContext

Episode

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// EpisodeNumber see : /episodeNumber
// Position of the episode within an ordered group of episodes.
EpisodeNumber interface{} `json:"episodeNumber"` // types : Integer Text

// MusicBy see : /musicBy
// The composer of the soundtrack.
MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

// PartOfSeason see : /partOfSeason
// The season to which this episode belongs.
PartOfSeason *CreativeWorkSeason `json:"partOfSeason"`

// PartOfSeries see : /partOfSeries
// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
PartOfSeries *CreativeWorkSeries `json:"partOfSeries"`

// ProductionCompany see : /productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc.
ProductionCompany *Organization `json:"productionCompany"`

// Trailer see : /trailer
// The trailer of a movie or tv/radio series, season, episode, etc.
Trailer *VideoObject `json:"trailer"`

}

func (v *RadioEpisode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioEpisode"

	return json.Marshal(v)
}
