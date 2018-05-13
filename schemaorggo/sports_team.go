package schemaorggo

import "encoding/json"

// SportsTeam see : https://schema.org/SportsTeam
type SportsTeam struct {
	SportsOrganization

	typeContext

	// Athlete see : https://schema.org/athlete
	// A person that acts as performing member of a sports team; a player as opposed to a coach.
	// types : Person
	Athlete []*Person `json:"athlete,omitempty"`

	// Coach see : https://schema.org/coach
	// A person that acts in a coaching role for a sports team.
	// types : Person
	Coach []*Person `json:"coach,omitempty"`
}

func (v SportsTeam) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.SportsOrganization.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Athlete
		if len(v.Athlete) == 1 {
			value = v.Athlete[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["athlete"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Coach
		if len(v.Coach) == 1 {
			value = v.Coach[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["coach"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SportsTeam) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SportsTeam"

	return data, nil
}

func (v SportsTeam) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
