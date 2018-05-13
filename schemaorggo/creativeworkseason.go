package schemaorggo

import "encoding/json"

// CreativeWorkSeason see : https://schema.org/CreativeWorkSeason
type CreativeWorkSeason struct {
	CreativeWork

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate []interface{} `json:"endDate,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	// types : Episode
	Episode []*Episode `json:"episode,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes []float64 `json:"numberOfEpisodes,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	// types : CreativeWorkSeries
	PartOfSeries []*CreativeWorkSeries `json:"partOfSeries,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany []*Organization `json:"productionCompany,omitempty"`

	// SeasonNumber see : https://schema.org/seasonNumber
	// Position of the season within an ordered group of seasons.
	// types : Integer Text
	SeasonNumber []interface{} `json:"seasonNumber,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate []interface{} `json:"startDate,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer []*VideoObject `json:"trailer,omitempty"`
}

func (v CreativeWorkSeason) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

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
		var value interface{} = v.EndDate
		if len(v.EndDate) == 1 {
			value = v.EndDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endDate"] = json.RawMessage(b)
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
		var value interface{} = v.SeasonNumber
		if len(v.SeasonNumber) == 1 {
			value = v.SeasonNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seasonNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StartDate
		if len(v.StartDate) == 1 {
			value = v.StartDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["startDate"] = json.RawMessage(b)
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

func (v CreativeWorkSeason) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CreativeWorkSeason"

	return data, nil
}

func (v CreativeWorkSeason) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
