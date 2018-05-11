package schemaorg

import "encoding/json"

// SportsTeam see : https://schema.org/SportsTeam
type SportsTeam struct {

typeContext

SportsOrganization

// Athlete see : https://schema.org/athlete
// A person that acts as performing member of a sports team; a player as opposed to a coach.
Athlete *Person `json:"athlete"`

// Coach see : https://schema.org/coach
// A person that acts in a coaching role for a sports team.
Coach *Person `json:"coach"`

}

func (v *SportsTeam) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SportsTeam"

	return json.Marshal(v)
}
