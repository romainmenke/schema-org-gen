package schemaorggo

import "encoding/json"

// MusicGroup see : https://schema.org/MusicGroup
type MusicGroup struct {
	PerformingGroup

	typeContext

	// Album see : https://schema.org/album
	// A music album. Supersedes albums (see: https://schema.org/albums).
	// types : MusicAlbum
	Album []*MusicAlbum `json:"album,omitempty"`

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	// types : Text URL
	Genre []string `json:"genre,omitempty"`

	// Track see : https://schema.org/track
	// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
	// types : ItemList MusicRecording
	Track []interface{} `json:"track,omitempty"`
}

func (v MusicGroup) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PerformingGroup.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Album
		if len(v.Album) == 1 {
			value = v.Album[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["album"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Genre
		if len(v.Genre) == 1 {
			value = v.Genre[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["genre"] = json.RawMessage(b)
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

func (v MusicGroup) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MusicGroup"

	return data, nil
}

func (v MusicGroup) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
