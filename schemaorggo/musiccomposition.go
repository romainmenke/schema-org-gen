package schemaorggo

import "encoding/json"

// MusicComposition see : https://schema.org/MusicComposition
type MusicComposition struct {
	CreativeWork

	typeContext

	// Composer see : https://schema.org/composer
	// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	Composer interface{} `json:"composer,omitempty"` // types : Organization Person

	// FirstPerformance see : https://schema.org/firstPerformance
	// The date and place the work was first performed.
	FirstPerformance *Event `json:"firstPerformance,omitempty"`

	// IncludedComposition see : https://schema.org/includedComposition
	// Smaller compositions included in this work (e.g. a movement in a symphony).
	IncludedComposition *MusicComposition `json:"includedComposition,omitempty"`

	// IswcCode see : https://schema.org/iswcCode
	// The International Standard Musical Work Code for the composition.
	IswcCode string `json:"iswcCode,omitempty"`

	// Lyricist see : https://schema.org/lyricist
	// The person who wrote the words.
	Lyricist *Person `json:"lyricist,omitempty"`

	// Lyrics see : https://schema.org/lyrics
	// The words in the song.
	Lyrics *CreativeWork `json:"lyrics,omitempty"`

	// MusicArrangement see : https://schema.org/musicArrangement
	// An arrangement derived from the composition.
	MusicArrangement *MusicComposition `json:"musicArrangement,omitempty"`

	// MusicCompositionForm see : https://schema.org/musicCompositionForm
	// The type of composition (e.g. overture, sonata, symphony, etc.).
	MusicCompositionForm string `json:"musicCompositionForm,omitempty"`

	// MusicalKey see : https://schema.org/musicalKey
	// The key, mode, or scale this composition uses.
	MusicalKey string `json:"musicalKey,omitempty"`

	// RecordedAs see : https://schema.org/recordedAs
	// An audio recording of the work. Inverse property: recordingOf (see: https://schema.org/recordingOf).
	RecordedAs *MusicRecording `json:"recordedAs,omitempty"`
}

func (v MusicComposition) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicComposition"

	return json.Marshal(v)
}

func (v *MusicComposition) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicComposition"

	return json.Marshal(*v)
}
