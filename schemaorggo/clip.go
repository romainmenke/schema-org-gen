package schemaorggo

import "encoding/json"

// Clip see : https://schema.org/Clip
type Clip struct {
	CreativeWork

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor"`

	// ClipNumber see : https://schema.org/clipNumber
	// Position of the clip within an ordered group of clips.
	ClipNumber interface{} `json:"clipNumber"` // types : Integer Text

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

	// PartOfEpisode see : https://schema.org/partOfEpisode
	// The episode to which this clip belongs.
	PartOfEpisode *Episode `json:"partOfEpisode"`

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries"`
}

func (v *Clip) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Clip"

	return json.Marshal(v)
}
