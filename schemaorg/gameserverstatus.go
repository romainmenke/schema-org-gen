package schemaorg

import "encoding/json"

// GameServerStatus see : https://schema.org/GameServerStatus
type GameServerStatus struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *GameServerStatus) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GameServerStatus"

	return json.Marshal(v)
}
