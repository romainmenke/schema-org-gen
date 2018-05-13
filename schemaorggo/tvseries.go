package schemaorggo

import "encoding/json"

// TVSeries see : https://schema.org/TVSeries
type TVSeries struct {
	CreativeWork

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// ContainsSeason see : https://schema.org/containsSeason
	// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
	// types : CreativeWorkSeason
	ContainsSeason []*CreativeWorkSeason `json:"containsSeason,omitempty"`

	// CountryOfOrigin see : https://schema.org/countryOfOrigin
	// The country of the principal offices of the production company or individual responsible for the movie or program.
	// types : Country
	CountryOfOrigin []*Country `json:"countryOfOrigin,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	// types : Episode
	Episode []*Episode `json:"episode,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy []interface{} `json:"musicBy,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes []float64 `json:"numberOfEpisodes,omitempty"`

	// NumberOfSeasons see : https://schema.org/numberOfSeasons
	// The number of seasons in this series.
	// types : Integer
	NumberOfSeasons []float64 `json:"numberOfSeasons,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany []*Organization `json:"productionCompany,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer []*VideoObject `json:"trailer,omitempty"`
}

func (v TVSeries) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.IntoMap(intop)

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
		var value interface{} = v.ContainsSeason
		if len(v.ContainsSeason) == 1 {
			value = v.ContainsSeason[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["containsSeason"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CountryOfOrigin
		if len(v.CountryOfOrigin) == 1 {
			value = v.CountryOfOrigin[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["countryOfOrigin"] = json.RawMessage(b)
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
		var value interface{} = v.Episode
		if len(v.Episode) == 1 {
			value = v.Episode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["episode"] = json.RawMessage(b)
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
		var value interface{} = v.NumberOfEpisodes
		if len(v.NumberOfEpisodes) == 1 {
			value = v.NumberOfEpisodes[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfEpisodes"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfSeasons
		if len(v.NumberOfSeasons) == 1 {
			value = v.NumberOfSeasons[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfSeasons"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductionCompany
		if len(v.ProductionCompany) == 1 {
			value = v.ProductionCompany[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productionCompany"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Trailer
		if len(v.Trailer) == 1 {
			value = v.Trailer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trailer"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TVSeries) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TVSeries"

	return data, nil
}

func (v TVSeries) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
