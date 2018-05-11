package schemaorg

import "encoding/json"

// MusicRelease see : https://schema.org/MusicRelease
type MusicRelease struct {

typeContext

MusicPlaylist

// CatalogNumber see : /catalogNumber
// The catalog number for the release.
CatalogNumber string `json:"catalogNumber"`

// CreditedTo see : /creditedTo
// The group the release is credited to if different than the byArtist. For example, Red and Blue is credited to "Stefani Germanotta Band", but by Lady Gaga.
CreditedTo interface{} `json:"creditedTo"` // types : Organization Person

// Duration see : /duration
// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
Duration *Duration `json:"duration"`

// MusicReleaseFormat see : /musicReleaseFormat
// Format of this release (the type of recording media used, ie. compact disc, digital media, LP, etc.).
MusicReleaseFormat *MusicReleaseFormatType `json:"musicReleaseFormat"`

// RecordLabel see : /recordLabel
// The label that issued the release.
RecordLabel *Organization `json:"recordLabel"`

// ReleaseOf see : /releaseOf
// The album this is a release of. Inverse property: albumRelease (see: https://schema.org/albumRelease).
ReleaseOf *MusicAlbum `json:"releaseOf"`

}

func (v *MusicRelease) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicRelease"

	return json.Marshal(v)
}
