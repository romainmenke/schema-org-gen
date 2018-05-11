package schemaorg

import "encoding/json"

// MusicPlaylist see : https://schema.org/MusicPlaylist
type MusicPlaylist struct {

typeContext

CreativeWork

// NumTracks see : https://schema.org/numTracks
// The number of tracks in this album or playlist.
NumTracks int `json:"numTracks"`

// Track see : https://schema.org/track
// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
Track interface{} `json:"track"` // types : ItemList MusicRecording

}

func (v *MusicPlaylist) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicPlaylist"

	return json.Marshal(v)
}
