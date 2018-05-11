package schemaorg

import "encoding/json"

// TechArticle see : https://schema.org/TechArticle
type TechArticle struct {

typeContext

Article

// Dependencies see : /dependencies
// Prerequisites needed to fulfill steps in article.
Dependencies string `json:"dependencies"`

// ProficiencyLevel see : /proficiencyLevel
// Proficiency needed for this content; expected values: 'Beginner', 'Expert'.
ProficiencyLevel string `json:"proficiencyLevel"`

}

func (v *TechArticle) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TechArticle"

	return json.Marshal(v)
}
