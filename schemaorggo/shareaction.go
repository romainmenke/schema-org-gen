package schemaorggo

import "encoding/json"

// ShareAction see : https://schema.org/ShareAction
type ShareAction struct {
	CommunicateAction

	typeContext

	// About see : https://schema.org/about
	// The subject matter of the content.
	// types : Thing
	About []*Thing `json:"about,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage []interface{} `json:"inLanguage,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient []interface{} `json:"recipient,omitempty"`
}

func (v ShareAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CommunicateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.About
		if len(v.About) == 1 {
			value = v.About[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["about"] = json.RawMessage(b)
		}
	}

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

	{
		var value interface{} = v.Recipient
		if len(v.Recipient) == 1 {
			value = v.Recipient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipient"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ShareAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ShareAction"

	return data, nil
}

func (v ShareAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
