package schemaorggo

import "encoding/json"

// MenuSection see : https://schema.org/MenuSection
type MenuSection struct {
	CreativeWork

	typeContext

	// HasMenuItem see : https://schema.org/hasMenuItem
	// A food or drink item contained in a menu or menu section.
	// types : MenuItem
	HasMenuItem []*MenuItem `json:"hasMenuItem,omitempty"`

	// HasMenuSection see : https://schema.org/hasMenuSection
	// A subgrouping of the menu (by dishes, course, serving time period, etc.).
	// types : MenuSection
	HasMenuSection []*MenuSection `json:"hasMenuSection,omitempty"`
}

func (v MenuSection) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.HasMenuItem
		if len(v.HasMenuItem) == 1 {
			value = v.HasMenuItem[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasMenuItem"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasMenuSection
		if len(v.HasMenuSection) == 1 {
			value = v.HasMenuSection[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasMenuSection"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MenuSection) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MenuSection"

	return data, nil
}

func (v MenuSection) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
