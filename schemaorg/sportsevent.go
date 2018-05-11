package schemaorg

import "encoding/json"

// SportsEvent see : https://schema.org/SportsEvent
type SportsEvent struct {

typeContext

Event

// AwayTeam see : /awayTeam
// The away team in a sports event.
AwayTeam interface{} `json:"awayTeam"` // types : Person SportsTeam

// Competitor see : /competitor
// A competitor in a sports event.
Competitor interface{} `json:"competitor"` // types : Person SportsTeam

// HomeTeam see : /homeTeam
// The home team in a sports event.
HomeTeam interface{} `json:"homeTeam"` // types : Person SportsTeam

}

func (v *SportsEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SportsEvent"

	return json.Marshal(v)
}
