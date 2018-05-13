package schemaorggo

import "encoding/json"

// MusicRecording see : https://schema.org/MusicRecording
type MusicRecording struct {
	CreativeWork

	typeContext

	// ByArtist see : https://schema.org/byArtist
	// The artist that performed this album or recording.
	// types : MusicGroup
	ByArtist []*MusicGroup `json:"byArtist,omitempty"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	Duration []*Duration `json:"duration,omitempty"`

	// InAlbum see : https://schema.org/inAlbum
	// The album to which this recording belongs.
	// types : MusicAlbum
	InAlbum []*MusicAlbum `json:"inAlbum,omitempty"`

	// InPlaylist see : https://schema.org/inPlaylist
	// The playlist to which this recording belongs.
	// types : MusicPlaylist
	InPlaylist []*MusicPlaylist `json:"inPlaylist,omitempty"`

	// IsrcCode see : https://schema.org/isrcCode
	// The International Standard Recording Code for the recording.
	// types : Text
	IsrcCode []string `json:"isrcCode,omitempty"`

	// RecordingOf see : https://schema.org/recordingOf
	// The composition this track is a recording of. Inverse property: recordedAs (see: https://schema.org/recordedAs).
	// types : MusicComposition
	RecordingOf []*MusicComposition `json:"recordingOf,omitempty"`
}

func (v MusicRecording) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ByArtist
		if len(v.ByArtist) == 1 {
			value = v.ByArtist[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["byArtist"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Duration
		if len(v.Duration) == 1 {
			value = v.Duration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["duration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InAlbum
		if len(v.InAlbum) == 1 {
			value = v.InAlbum[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inAlbum"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InPlaylist
		if len(v.InPlaylist) == 1 {
			value = v.InPlaylist[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inPlaylist"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsrcCode
		if len(v.IsrcCode) == 1 {
			value = v.IsrcCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isrcCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecordingOf
		if len(v.RecordingOf) == 1 {
			value = v.RecordingOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recordingOf"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MusicRecording) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MusicRecording"

	return data, nil
}

func (v MusicRecording) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
