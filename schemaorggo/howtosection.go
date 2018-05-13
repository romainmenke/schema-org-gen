package schemaorggo

import "encoding/json"

// HowToSection see : https://schema.org/HowToSection
type HowToSection struct {
	ItemList

	typeContext

	// Steps see : https://schema.org/steps
	// The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	// types : CreativeWork ItemList Text
	Steps []interface{} `json:"steps,omitempty"`
}

func (v HowToSection) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ItemList.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Steps
		if len(v.Steps) == 1 {
			value = v.Steps[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["steps"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v HowToSection) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowToSection"

	return data, nil
}

func (v HowToSection) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
