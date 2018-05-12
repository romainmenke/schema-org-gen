package schemaorggo

import "encoding/json"

// RadioClip see : https://schema.org/RadioClip
type RadioClip struct {
	Clip

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor,omitempty"`

	// ClipNumber see : https://schema.org/clipNumber
	// Position of the clip within an ordered group of clips.
	ClipNumber interface{} `json:"clipNumber,omitempty"` // types : Integer Text

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy,omitempty"` // types : MusicGroup Person

	// PartOfEpisode see : https://schema.org/partOfEpisode
	// The episode to which this clip belongs.
	PartOfEpisode *Episode `json:"partOfEpisode,omitempty"`

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty"`
}

func (v RadioClip) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RadioClip"

	return json.Marshal(v)
}

func (v *RadioClip) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RadioClip"

	return json.Marshal(*v)
}
