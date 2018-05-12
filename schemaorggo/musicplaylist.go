package schemaorggo

import "encoding/json"

// MusicPlaylist see : https://schema.org/MusicPlaylist
type MusicPlaylist struct {
	CreativeWork

	typeContext

	// NumTracks see : https://schema.org/numTracks
	// The number of tracks in this album or playlist.
	NumTracks float64 `json:"numTracks,omitempty"` // types : Integer

	// Track see : https://schema.org/track
	// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
	Track interface{} `json:"track,omitempty"` // types : ItemList MusicRecording

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
