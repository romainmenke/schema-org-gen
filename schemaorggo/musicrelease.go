package schemaorggo

import "encoding/json"

// MusicRelease see : https://schema.org/MusicRelease
type MusicRelease struct {
	MusicPlaylist

	typeContext

	// CatalogNumber see : https://schema.org/catalogNumber
	// The catalog number for the release.
	// types : Text
	CatalogNumber string `json:"catalogNumber,omitempty"`

	// CreditedTo see : https://schema.org/creditedTo
	// The group the release is credited to if different than the byArtist. For example, Red and Blue is credited to &quot;Stefani Germanotta Band&quot;, but by Lady Gaga.
	// types : Organization Person
	CreditedTo interface{} `json:"creditedTo,omitempty"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	Duration *Duration `json:"duration,omitempty"`

	// MusicReleaseFormat see : https://schema.org/musicReleaseFormat
	// Format of this release (the type of recording media used, ie. compact disc, digital media, LP, etc.).
	// types : MusicReleaseFormatType
	MusicReleaseFormat *MusicReleaseFormatType `json:"musicReleaseFormat,omitempty"`

	// RecordLabel see : https://schema.org/recordLabel
	// The label that issued the release.
	// types : Organization
	RecordLabel *Organization `json:"recordLabel,omitempty"`

	// ReleaseOf see : https://schema.org/releaseOf
	// The album this is a release of. Inverse property: albumRelease (see: https://schema.org/albumRelease).
	// types : MusicAlbum
	ReleaseOf *MusicAlbum `json:"releaseOf,omitempty"`
}

func (v MusicRelease) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicRelease"

	return json.Marshal(v)
}

func (v *MusicRelease) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicRelease"

	return json.Marshal(*v)
}
