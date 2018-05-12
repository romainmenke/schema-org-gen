package schemaorg

import "encoding/json"

// MusicAlbum see : https://schema.org/MusicAlbum
type MusicAlbum struct {

MusicPlaylist

typeContext

// AlbumProductionType see : https://schema.org/albumProductionType
// Classification of the album by it's type of content: soundtrack, live album, studio album, etc.
AlbumProductionType *MusicAlbumProductionType `json:"albumProductionType"`

// AlbumRelease see : https://schema.org/albumRelease
// A release of this album. Inverse property: releaseOf (see: https://schema.org/releaseOf).
AlbumRelease *MusicRelease `json:"albumRelease"`

// AlbumReleaseType see : https://schema.org/albumReleaseType
// The kind of release which this album is: single, EP or album.
AlbumReleaseType *MusicAlbumReleaseType `json:"albumReleaseType"`

// ByArtist see : https://schema.org/byArtist
// The artist that performed this album or recording.
ByArtist *MusicGroup `json:"byArtist"`

}

func (v *MusicAlbum) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicAlbum"

	return json.Marshal(v)
}
