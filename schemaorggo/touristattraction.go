package schemaorggo

import "encoding/json"

// TouristAttraction see : https://schema.org/TouristAttraction
type TouristAttraction struct {
	Place

	typeContext

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	// types : Language Text
	AvailableLanguage []interface{} `json:"availableLanguage,omitempty"`

	// TouristType see : https://schema.org/touristType
	// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
	// types : Audience Text
	TouristType []interface{} `json:"touristType,omitempty"`
}

func (v TouristAttraction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Place.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AvailableLanguage
		if len(v.AvailableLanguage) == 1 {
			value = v.AvailableLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TouristType
		if len(v.TouristType) == 1 {
			value = v.TouristType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["touristType"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TouristAttraction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TouristAttraction"

	return data, nil
}

func (v TouristAttraction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
