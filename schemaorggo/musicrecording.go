package schemaorggo

import "encoding/json"

// MusicRecording see : https://schema.org/MusicRecording
type MusicRecording struct {
	CreativeWork

	typeContext

	// ByArtist see : https://schema.org/byArtist
	// The artist that performed this album or recording.
	ByArtist *MusicGroup `json:"byArtist"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	Duration *Duration `json:"duration"`

	// InAlbum see : https://schema.org/inAlbum
	// The album to which this recording belongs.
	InAlbum *MusicAlbum `json:"inAlbum"`

	// InPlaylist see : https://schema.org/inPlaylist
	// The playlist to which this recording belongs.
	InPlaylist *MusicPlaylist `json:"inPlaylist"`

	// IsrcCode see : https://schema.org/isrcCode
	// The International Standard Recording Code for the recording.
	IsrcCode string `json:"isrcCode"`

	// RecordingOf see : https://schema.org/recordingOf
	// The composition this track is a recording of. Inverse property: recordedAs (see: https://schema.org/recordedAs).
	RecordingOf *MusicComposition `json:"recordingOf"`
}

func (v *MusicRecording) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicRecording"

	return json.Marshal(v)
}
