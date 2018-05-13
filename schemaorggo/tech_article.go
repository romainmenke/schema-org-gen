package schemaorggo

import "encoding/json"

// TechArticle see : https://schema.org/TechArticle
type TechArticle struct {
	Article

	typeContext

	// Dependencies see : https://schema.org/dependencies
	// Prerequisites needed to fulfill steps in article.
	// types : Text
	Dependencies []string `json:"dependencies,omitempty"`

	// ProficiencyLevel see : https://schema.org/proficiencyLevel
	// Proficiency needed for this content; expected values: &#39;Beginner&#39;, &#39;Expert&#39;.
	// types : Text
	ProficiencyLevel []string `json:"proficiencyLevel,omitempty"`
}

func (v TechArticle) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Article.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Dependencies
		if len(v.Dependencies) == 1 {
			value = v.Dependencies[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dependencies"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProficiencyLevel
		if len(v.ProficiencyLevel) == 1 {
			value = v.ProficiencyLevel[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["proficiencyLevel"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TechArticle) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TechArticle"

	return data, nil
}

func (v TechArticle) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
