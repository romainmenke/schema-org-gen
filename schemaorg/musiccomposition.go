package schemaorg

import "encoding/json"

// MusicComposition see : https://schema.org/MusicComposition
type MusicComposition struct {

typeContext

CreativeWork

// Composer see : https://schema.org/composer
// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
Composer interface{} `json:"composer"` // types : Organization Person

// FirstPerformance see : https://schema.org/firstPerformance
// The date and place the work was first performed.
FirstPerformance *Event `json:"firstPerformance"`

// IncludedComposition see : https://schema.org/includedComposition
// Smaller compositions included in this work (e.g. a movement in a symphony).
IncludedComposition *MusicComposition `json:"includedComposition"`

// IswcCode see : https://schema.org/iswcCode
// The International Standard Musical Work Code for the composition.
IswcCode string `json:"iswcCode"`

// Lyricist see : https://schema.org/lyricist
// The person who wrote the words.
Lyricist *Person `json:"lyricist"`

// Lyrics see : https://schema.org/lyrics
// The words in the song.
Lyrics *CreativeWork `json:"lyrics"`

// MusicArrangement see : https://schema.org/musicArrangement
// An arrangement derived from the composition.
MusicArrangement *MusicComposition `json:"musicArrangement"`

// MusicCompositionForm see : https://schema.org/musicCompositionForm
// The type of composition (e.g. overture, sonata, symphony, etc.).
MusicCompositionForm string `json:"musicCompositionForm"`

// MusicalKey see : https://schema.org/musicalKey
// The key, mode, or scale this composition uses.
MusicalKey string `json:"musicalKey"`

// RecordedAs see : https://schema.org/recordedAs
// An audio recording of the work. Inverse property: recordingOf (see: https://schema.org/recordingOf).
RecordedAs *MusicRecording `json:"recordedAs"`

}

func (v *MusicComposition) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicComposition"

	return json.Marshal(v)
}
