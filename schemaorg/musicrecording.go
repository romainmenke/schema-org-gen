package schemaorg

import "encoding/json"

// MusicRecording see : https://schema.org/MusicRecording
type MusicRecording struct {

typeContext

CreativeWork

// ByArtist see : /byArtist
// The artist that performed this album or recording.
ByArtist *MusicGroup `json:"byArtist"`

// Duration see : /duration
// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
Duration *Duration `json:"duration"`

// InAlbum see : /inAlbum
// The album to which this recording belongs.
InAlbum *MusicAlbum `json:"inAlbum"`

// InPlaylist see : /inPlaylist
// The playlist to which this recording belongs.
InPlaylist *MusicPlaylist `json:"inPlaylist"`

// IsrcCode see : /isrcCode
// The International Standard Recording Code for the recording.
IsrcCode string `json:"isrcCode"`

// RecordingOf see : /recordingOf
// The composition this track is a recording of. Inverse property: recordedAs (see: https://schema.org/recordedAs).
RecordingOf *MusicComposition `json:"recordingOf"`

}

func (v *MusicRecording) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicRecording"

	return json.Marshal(v)
}
