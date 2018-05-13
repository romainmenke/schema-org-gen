package schemaorggo

import "encoding/json"

// MusicAlbum see : https://schema.org/MusicAlbum
type MusicAlbum struct {
	MusicPlaylist

	typeContext

	// AlbumProductionType see : https://schema.org/albumProductionType
	// Classification of the album by it&#39;s type of content: soundtrack, live album, studio album, etc.
	// types : MusicAlbumProductionType
	AlbumProductionType []*MusicAlbumProductionType `json:"albumProductionType,omitempty"`

	// AlbumRelease see : https://schema.org/albumRelease
	// A release of this album. Inverse property: releaseOf (see: https://schema.org/releaseOf).
	// types : MusicRelease
	AlbumRelease []*MusicRelease `json:"albumRelease,omitempty"`

	// AlbumReleaseType see : https://schema.org/albumReleaseType
	// The kind of release which this album is: single, EP or album.
	// types : MusicAlbumReleaseType
	AlbumReleaseType []*MusicAlbumReleaseType `json:"albumReleaseType,omitempty"`

	// ByArtist see : https://schema.org/byArtist
	// The artist that performed this album or recording.
	// types : MusicGroup
	ByArtist []*MusicGroup `json:"byArtist,omitempty"`
}

func (v MusicAlbum) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.MusicPlaylist.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AlbumProductionType
		if len(v.AlbumProductionType) == 1 {
			value = v.AlbumProductionType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["albumProductionType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AlbumRelease
		if len(v.AlbumRelease) == 1 {
			value = v.AlbumRelease[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["albumRelease"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AlbumReleaseType
		if len(v.AlbumReleaseType) == 1 {
			value = v.AlbumReleaseType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["albumReleaseType"] = json.RawMessage(b)
		}
	}

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

	*intop = into

	return nil
}

func (v MusicAlbum) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MusicAlbum"

	return data, nil
}

func (v MusicAlbum) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
