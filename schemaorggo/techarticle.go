package schemaorggo

import "encoding/json"

// TechArticle see : https://schema.org/TechArticle
type TechArticle struct {
	Article

	typeContext

	// Dependencies see : https://schema.org/dependencies
	// Prerequisites needed to fulfill steps in article.
	// types : Text
	Dependencies string `json:"dependencies,omitempty"`

	// ProficiencyLevel see : https://schema.org/proficiencyLevel
	// Proficiency needed for this content; expected values: &#39;Beginner&#39;, &#39;Expert&#39;.
	// types : Text
	ProficiencyLevel string `json:"proficiencyLevel,omitempty"`
}

func (v TechArticle) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TechArticle"

	return json.Marshal(v)
}

func (v *TechArticle) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TechArticle"

	return json.Marshal(*v)
}
