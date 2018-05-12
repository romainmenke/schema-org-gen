package schemaorggo

import "encoding/json"

// SportsEvent see : https://schema.org/SportsEvent
type SportsEvent struct {
	Event

	typeContext

	// AwayTeam see : https://schema.org/awayTeam
	// The away team in a sports event.
	AwayTeam interface{} `json:"awayTeam,omitempty"` // types : Person SportsTeam

	// Competitor see : https://schema.org/competitor
	// A competitor in a sports event.
	Competitor interface{} `json:"competitor,omitempty"` // types : Person SportsTeam

	// HomeTeam see : https://schema.org/homeTeam
	// The home team in a sports event.
	HomeTeam interface{} `json:"homeTeam,omitempty"` // types : Person SportsTeam

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
