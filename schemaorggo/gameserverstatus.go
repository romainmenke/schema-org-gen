package schemaorggo

import "encoding/json"

// GameServerStatus see : https://schema.org/GameServerStatus
type GameServerStatus struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v GameServerStatus) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GameServerStatus"

	return json.Marshal(v)
}

func (v *GameServerStatus) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GameServerStatus"

	return json.Marshal(*v)
}
