package schemaorggo

import "encoding/json"

// TVSeason see : https://schema.org/TVSeason
type TVSeason struct {
	CreativeWork

	typeContext

	// CountryOfOrigin see : https://schema.org/countryOfOrigin
	// The country of the principal offices of the production company or individual responsible for the movie or program.
	// types : Country
	CountryOfOrigin []*Country `json:"countryOfOrigin,omitempty"`
}

func (v TVSeason) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

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

	*intop = into

	return nil
}

func (v TVSeason) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TVSeason"

	return data, nil
}

func (v TVSeason) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
