package schemaorg

import "encoding/json"

// MusicAlbumProductionType see : https://schema.org/MusicAlbumProductionType
type MusicAlbumProductionType struct {

Enumeration

typeContext

// SupersededBy see : http://meta.google.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *MusicAlbumProductionType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicAlbumProductionType"

	return json.Marshal(v)
}
