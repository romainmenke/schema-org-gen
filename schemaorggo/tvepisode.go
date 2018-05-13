package schemaorggo

import "encoding/json"

// TVEpisode see : https://schema.org/TVEpisode
type TVEpisode struct {
	Episode

	typeContext

	// CountryOfOrigin see : https://schema.org/countryOfOrigin
	// The country of the principal offices of the production company or individual responsible for the movie or program.
	// types : Country
	CountryOfOrigin []*Country `json:"countryOfOrigin,omitempty"`

	// SubtitleLanguage see : https://schema.org/subtitleLanguage
	// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	// types : Language Text
	SubtitleLanguage []interface{} `json:"subtitleLanguage,omitempty"`
}

func (v TVEpisode) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Episode.intoMap(intop)

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

	{
		var value interface{} = v.SubtitleLanguage
		if len(v.SubtitleLanguage) == 1 {
			value = v.SubtitleLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subtitleLanguage"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TVEpisode) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TVEpisode"

	return data, nil
}

func (v TVEpisode) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
