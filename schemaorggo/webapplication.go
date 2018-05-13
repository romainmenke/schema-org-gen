package schemaorggo

import "encoding/json"

// WebApplication see : https://schema.org/WebApplication
type WebApplication struct {
	SoftwareApplication

	typeContext

	// BrowserRequirements see : https://schema.org/browserRequirements
	// Specifies browser requirements in human-readable text. For example, &#39;requires HTML5 support&#39;.
	// types : Text
	BrowserRequirements []string `json:"browserRequirements,omitempty"`
}

func (v WebApplication) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.SoftwareApplication.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BrowserRequirements
		if len(v.BrowserRequirements) == 1 {
			value = v.BrowserRequirements[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["browserRequirements"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v WebApplication) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WebApplication"

	return data, nil
}

func (v WebApplication) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
