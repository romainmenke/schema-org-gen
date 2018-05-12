package schemaorg

import "encoding/json"

// MusicGroup see : https://schema.org/MusicGroup
type MusicGroup struct {

PerformingGroup

typeContext

// Album see : https://schema.org/album
// A music album. Supersedes albums (see: https://schema.org/albums).
Album *MusicAlbum `json:"album"`

// Genre see : https://schema.org/genre
// Genre of the creative work, broadcast channel or group.
Genre interface{} `json:"genre"` // types : Text URL

// Track see : https://schema.org/track
// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
Track interface{} `json:"track"` // types : ItemList MusicRecording

}

func (v *MusicGroup) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicGroup"

	return json.Marshal(v)
}
