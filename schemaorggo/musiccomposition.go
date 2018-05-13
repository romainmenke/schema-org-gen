package schemaorggo

import "encoding/json"

// MusicComposition see : https://schema.org/MusicComposition
type MusicComposition struct {
	CreativeWork

	typeContext

	// Composer see : https://schema.org/composer
	// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	// types : Organization Person
	Composer []interface{} `json:"composer,omitempty"`

	// FirstPerformance see : https://schema.org/firstPerformance
	// The date and place the work was first performed.
	// types : Event
	FirstPerformance []*Event `json:"firstPerformance,omitempty"`

	// IncludedComposition see : https://schema.org/includedComposition
	// Smaller compositions included in this work (e.g. a movement in a symphony).
	// types : MusicComposition
	IncludedComposition []*MusicComposition `json:"includedComposition,omitempty"`

	// IswcCode see : https://schema.org/iswcCode
	// The International Standard Musical Work Code for the composition.
	// types : Text
	IswcCode []string `json:"iswcCode,omitempty"`

	// Lyricist see : https://schema.org/lyricist
	// The person who wrote the words.
	// types : Person
	Lyricist []*Person `json:"lyricist,omitempty"`

	// Lyrics see : https://schema.org/lyrics
	// The words in the song.
	// types : CreativeWork
	Lyrics []*CreativeWork `json:"lyrics,omitempty"`

	// MusicArrangement see : https://schema.org/musicArrangement
	// An arrangement derived from the composition.
	// types : MusicComposition
	MusicArrangement []*MusicComposition `json:"musicArrangement,omitempty"`

	// MusicCompositionForm see : https://schema.org/musicCompositionForm
	// The type of composition (e.g. overture, sonata, symphony, etc.).
	// types : Text
	MusicCompositionForm []string `json:"musicCompositionForm,omitempty"`

	// MusicalKey see : https://schema.org/musicalKey
	// The key, mode, or scale this composition uses.
	// types : Text
	MusicalKey []string `json:"musicalKey,omitempty"`

	// RecordedAs see : https://schema.org/recordedAs
	// An audio recording of the work. Inverse property: recordingOf (see: https://schema.org/recordingOf).
	// types : MusicRecording
	RecordedAs []*MusicRecording `json:"recordedAs,omitempty"`
}

func (v MusicComposition) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Composer
		if len(v.Composer) == 1 {
			value = v.Composer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["composer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FirstPerformance
		if len(v.FirstPerformance) == 1 {
			value = v.FirstPerformance[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["firstPerformance"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IncludedComposition
		if len(v.IncludedComposition) == 1 {
			value = v.IncludedComposition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["includedComposition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IswcCode
		if len(v.IswcCode) == 1 {
			value = v.IswcCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["iswcCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Lyricist
		if len(v.Lyricist) == 1 {
			value = v.Lyricist[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lyricist"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Lyrics
		if len(v.Lyrics) == 1 {
			value = v.Lyrics[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lyrics"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicArrangement
		if len(v.MusicArrangement) == 1 {
			value = v.MusicArrangement[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicArrangement"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicCompositionForm
		if len(v.MusicCompositionForm) == 1 {
			value = v.MusicCompositionForm[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicCompositionForm"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicalKey
		if len(v.MusicalKey) == 1 {
			value = v.MusicalKey[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicalKey"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecordedAs
		if len(v.RecordedAs) == 1 {
			value = v.RecordedAs[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recordedAs"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MusicComposition) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MusicComposition"

	return data, nil
}

func (v MusicComposition) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
