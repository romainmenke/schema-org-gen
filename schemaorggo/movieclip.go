package schemaorggo

import "encoding/json"

// MovieClip see : https://schema.org/MovieClip
type MovieClip struct {
	Clip

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// ClipNumber see : https://schema.org/clipNumber
	// Position of the clip within an ordered group of clips.
	// types : Integer Text
	ClipNumber interface{} `json:"clipNumber,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy interface{} `json:"musicBy,omitempty"`

	// PartOfEpisode see : https://schema.org/partOfEpisode
	// The episode to which this clip belongs.
	// types : Episode
	PartOfEpisode *Episode `json:"partOfEpisode,omitempty"`

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	// types : CreativeWorkSeason
	PartOfSeason *CreativeWorkSeason `json:"partOfSeason,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	// types : CreativeWorkSeries
	PartOfSeries *CreativeWorkSeries `json:"partOfSeries,omitempty"`
}

func (v MovieClip) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MovieClip"

	return json.Marshal(v)
}

func (v *MovieClip) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MovieClip"

	return json.Marshal(*v)
}
