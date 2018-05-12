package schemaorggo

import "encoding/json"

// MusicAlbumReleaseType see : https://schema.org/MusicAlbumReleaseType
type MusicAlbumReleaseType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v MusicAlbumReleaseType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicAlbumReleaseType"

	return json.Marshal(v)
}

func (v *MusicAlbumReleaseType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicAlbumReleaseType"

	return json.Marshal(*v)
}
