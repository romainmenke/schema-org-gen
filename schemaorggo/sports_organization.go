package schemaorggo

import "encoding/json"

// SportsOrganization see : https://schema.org/SportsOrganization
type SportsOrganization struct {
	Organization

	typeContext

	// Sport see : https://schema.org/sport
	// A type of sport (e.g. Baseball).
	// types : Text URL
	Sport []string `json:"sport,omitempty"`
}

func (v SportsOrganization) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Organization.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Sport
		if len(v.Sport) == 1 {
			value = v.Sport[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sport"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SportsOrganization) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SportsOrganization"

	return data, nil
}

func (v SportsOrganization) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
