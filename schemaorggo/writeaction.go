package schemaorggo

import "encoding/json"

// WriteAction see : https://schema.org/WriteAction
type WriteAction struct {
	CreateAction

	typeContext

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage []interface{} `json:"inLanguage,omitempty"`
}

func (v WriteAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.InLanguage
		if len(v.InLanguage) == 1 {
			value = v.InLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inLanguage"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v WriteAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WriteAction"

	return data, nil
}

func (v WriteAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
