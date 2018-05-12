package schemaorggo

import "encoding/json"

// SportsTeam see : https://schema.org/SportsTeam
type SportsTeam struct {
	SportsOrganization

	typeContext

	// Athlete see : https://schema.org/athlete
	// A person that acts as performing member of a sports team; a player as opposed to a coach.
	// types : Person
	Athlete *Person `json:"athlete,omitempty"`

	// Coach see : https://schema.org/coach
	// A person that acts in a coaching role for a sports team.
	// types : Person
	Coach *Person `json:"coach,omitempty"`
}

func (v SportsTeam) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SportsTeam"

	return json.Marshal(v)
}

func (v *SportsTeam) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SportsTeam"

	return json.Marshal(*v)
}
