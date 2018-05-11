package schemaorg

import "encoding/json"

// GamePlayMode see : https://schema.org/GamePlayMode
type GamePlayMode struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *GamePlayMode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GamePlayMode"

	return json.Marshal(v)
}
