package schemaorggo

import "encoding/json"

// MusicRelease see : https://schema.org/MusicRelease
type MusicRelease struct {
	MusicPlaylist

	typeContext

	// CatalogNumber see : https://schema.org/catalogNumber
	// The catalog number for the release.
	// types : Text
	CatalogNumber []string `json:"catalogNumber,omitempty"`

	// CreditedTo see : https://schema.org/creditedTo
	// The group the release is credited to if different than the byArtist. For example, Red and Blue is credited to &quot;Stefani Germanotta Band&quot;, but by Lady Gaga.
	// types : Organization Person
	CreditedTo []interface{} `json:"creditedTo,omitempty"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	Duration []*Duration `json:"duration,omitempty"`

	// MusicReleaseFormat see : https://schema.org/musicReleaseFormat
	// Format of this release (the type of recording media used, ie. compact disc, digital media, LP, etc.).
	// types : MusicReleaseFormatType
	MusicReleaseFormat []*MusicReleaseFormatType `json:"musicReleaseFormat,omitempty"`

	// RecordLabel see : https://schema.org/recordLabel
	// The label that issued the release.
	// types : Organization
	RecordLabel []*Organization `json:"recordLabel,omitempty"`

	// ReleaseOf see : https://schema.org/releaseOf
	// The album this is a release of. Inverse property: albumRelease (see: https://schema.org/albumRelease).
	// types : MusicAlbum
	ReleaseOf []*MusicAlbum `json:"releaseOf,omitempty"`
}

func (v MusicRelease) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.MusicPlaylist.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CatalogNumber
		if len(v.CatalogNumber) == 1 {
			value = v.CatalogNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["catalogNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CreditedTo
		if len(v.CreditedTo) == 1 {
			value = v.CreditedTo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["creditedTo"] = json.RawMessage(b)
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
		var value interface{} = v.MusicReleaseFormat
		if len(v.MusicReleaseFormat) == 1 {
			value = v.MusicReleaseFormat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicReleaseFormat"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecordLabel
		if len(v.RecordLabel) == 1 {
			value = v.RecordLabel[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recordLabel"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReleaseOf
		if len(v.ReleaseOf) == 1 {
			value = v.ReleaseOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["releaseOf"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MusicRelease) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MusicRelease"

	return data, nil
}

func (v MusicRelease) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
