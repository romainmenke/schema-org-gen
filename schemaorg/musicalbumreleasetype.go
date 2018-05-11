package schemaorg

import "encoding/json"

// MusicAlbumReleaseType see : https://schema.org/MusicAlbumReleaseType
type MusicAlbumReleaseType struct {

typeContext

Enumeration

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *MusicAlbumReleaseType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicAlbumReleaseType"

	return json.Marshal(v)
}
