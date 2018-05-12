package schemaorggo

import "encoding/json"

// MusicAlbumProductionType see : https://schema.org/MusicAlbumProductionType
type MusicAlbumProductionType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.google.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v MusicAlbumProductionType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicAlbumProductionType"

	return json.Marshal(v)
}

func (v *MusicAlbumProductionType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicAlbumProductionType"

	return json.Marshal(*v)
}
