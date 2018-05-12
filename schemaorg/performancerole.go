package schemaorg

import "encoding/json"

// PerformanceRole see : https://schema.org/PerformanceRole
type PerformanceRole struct {

Role

typeContext

// CharacterName see : https://schema.org/characterName
// The name of a character played in some acting or performing role, i.e. in a PerformanceRole.
CharacterName string `json:"characterName"`

}

func (v *PerformanceRole) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PerformanceRole"

	return json.Marshal(v)
}
