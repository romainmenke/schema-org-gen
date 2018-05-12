package schemaorggo

import "encoding/json"

// MusicAlbum see : https://schema.org/MusicAlbum
type MusicAlbum struct {
	MusicPlaylist

	typeContext

	// AlbumProductionType see : https://schema.org/albumProductionType
	// Classification of the album by it&#39;s type of content: soundtrack, live album, studio album, etc.
	// types : MusicAlbumProductionType
	AlbumProductionType *MusicAlbumProductionType `json:"albumProductionType,omitempty"`

	// AlbumRelease see : https://schema.org/albumRelease
	// A release of this album. Inverse property: releaseOf (see: https://schema.org/releaseOf).
	// types : MusicRelease
	AlbumRelease *MusicRelease `json:"albumRelease,omitempty"`

	// AlbumReleaseType see : https://schema.org/albumReleaseType
	// The kind of release which this album is: single, EP or album.
	// types : MusicAlbumReleaseType
	AlbumReleaseType *MusicAlbumReleaseType `json:"albumReleaseType,omitempty"`

	// ByArtist see : https://schema.org/byArtist
	// The artist that performed this album or recording.
	// types : MusicGroup
	ByArtist *MusicGroup `json:"byArtist,omitempty"`
}

func (v MusicAlbum) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicAlbum"

	return json.Marshal(v)
}

func (v *MusicAlbum) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicAlbum"

	return json.Marshal(*v)
}
