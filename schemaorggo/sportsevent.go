package schemaorggo

import "encoding/json"

// SportsEvent see : https://schema.org/SportsEvent
type SportsEvent struct {
	Event

	typeContext

	// AwayTeam see : https://schema.org/awayTeam
	// The away team in a sports event.
	// types : Person SportsTeam
	AwayTeam interface{} `json:"awayTeam,omitempty"`

	// Competitor see : https://schema.org/competitor
	// A competitor in a sports event.
	// types : Person SportsTeam
	Competitor interface{} `json:"competitor,omitempty"`

	// HomeTeam see : https://schema.org/homeTeam
	// The home team in a sports event.
	// types : Person SportsTeam
	HomeTeam interface{} `json:"homeTeam,omitempty"`
}

func (v SportsEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SportsEvent"

	return json.Marshal(v)
}

func (v *SportsEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SportsEvent"

	return json.Marshal(*v)
}
