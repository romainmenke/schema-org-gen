package schemaorggo

import "encoding/json"

// SportsEvent see : https://schema.org/SportsEvent
type SportsEvent struct {
	Event

	typeContext

	// AwayTeam see : https://schema.org/awayTeam
	// The away team in a sports event.
	// types : Person SportsTeam
	AwayTeam []interface{} `json:"awayTeam,omitempty"`

	// Competitor see : https://schema.org/competitor
	// A competitor in a sports event.
	// types : Person SportsTeam
	Competitor []interface{} `json:"competitor,omitempty"`

	// HomeTeam see : https://schema.org/homeTeam
	// The home team in a sports event.
	// types : Person SportsTeam
	HomeTeam []interface{} `json:"homeTeam,omitempty"`
}

func (v SportsEvent) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Event.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AwayTeam
		if len(v.AwayTeam) == 1 {
			value = v.AwayTeam[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["awayTeam"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Competitor
		if len(v.Competitor) == 1 {
			value = v.Competitor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["competitor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HomeTeam
		if len(v.HomeTeam) == 1 {
			value = v.HomeTeam[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["homeTeam"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SportsEvent) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SportsEvent"

	return data, nil
}

func (v SportsEvent) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
