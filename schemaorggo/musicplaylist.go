package schemaorggo

import "encoding/json"

// MusicPlaylist see : https://schema.org/MusicPlaylist
type MusicPlaylist struct {
	CreativeWork

	typeContext

	// NumTracks see : https://schema.org/numTracks
	// The number of tracks in this album or playlist.
	// types : Integer
	NumTracks []float64 `json:"numTracks,omitempty"`

	// Track see : https://schema.org/track
	// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
	// types : ItemList MusicRecording
	Track []interface{} `json:"track,omitempty"`
}

func (v MusicPlaylist) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.NumTracks
		if len(v.NumTracks) == 1 {
			value = v.NumTracks[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numTracks"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Track
		if len(v.Track) == 1 {
			value = v.Track[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["track"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MusicPlaylist) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MusicPlaylist"

	return data, nil
}

func (v MusicPlaylist) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
