package schemaorggo

import "encoding/json"

// MusicPlaylist see : https://schema.org/MusicPlaylist
type MusicPlaylist struct {
	CreativeWork

	typeContext

	// NumTracks see : https://schema.org/numTracks
	// The number of tracks in this album or playlist.
	// types : Integer
	NumTracks float64 `json:"numTracks,omitempty"`

	// Track see : https://schema.org/track
	// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
	// types : ItemList MusicRecording
	Track interface{} `json:"track,omitempty"`
}

func (v MusicPlaylist) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicPlaylist"

	return json.Marshal(v)
}

func (v *MusicPlaylist) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicPlaylist"

	return json.Marshal(*v)
}
