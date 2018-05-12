package schemaorggo

import "encoding/json"

// GamePlayMode see : https://schema.org/GamePlayMode
type GamePlayMode struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v GamePlayMode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GamePlayMode"

	return json.Marshal(v)
}

func (v *GamePlayMode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GamePlayMode"

	return json.Marshal(*v)
}
