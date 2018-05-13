package schemaorggo

import "encoding/json"

// MovieClip see : https://schema.org/MovieClip
type MovieClip struct {
	Clip

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// ClipNumber see : https://schema.org/clipNumber
	// Position of the clip within an ordered group of clips.
	// types : Integer Text
	ClipNumber []interface{} `json:"clipNumber,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy []interface{} `json:"musicBy,omitempty"`

	// PartOfEpisode see : https://schema.org/partOfEpisode
	// The episode to which this clip belongs.
	// types : Episode
	PartOfEpisode []*Episode `json:"partOfEpisode,omitempty"`

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	// types : CreativeWorkSeason
	PartOfSeason []*CreativeWorkSeason `json:"partOfSeason,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	// types : CreativeWorkSeries
	PartOfSeries []*CreativeWorkSeries `json:"partOfSeries,omitempty"`
}

func (v MovieClip) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Clip.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Actor
		if len(v.Actor) == 1 {
			value = v.Actor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ClipNumber
		if len(v.ClipNumber) == 1 {
			value = v.ClipNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["clipNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Director
		if len(v.Director) == 1 {
			value = v.Director[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["director"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicBy
		if len(v.MusicBy) == 1 {
			value = v.MusicBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfEpisode
		if len(v.PartOfEpisode) == 1 {
			value = v.PartOfEpisode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfEpisode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfSeason
		if len(v.PartOfSeason) == 1 {
			value = v.PartOfSeason[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfSeason"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfSeries
		if len(v.PartOfSeries) == 1 {
			value = v.PartOfSeries[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfSeries"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MovieClip) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MovieClip"

	return data, nil
}

func (v MovieClip) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
