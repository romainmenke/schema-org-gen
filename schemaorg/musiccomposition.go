package schemaorg

import "encoding/json"

// MusicComposition see : https://schema.org/MusicComposition
type MusicComposition struct {

typeContext

CreativeWork

// Composer see : /composer
// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
Composer interface{} `json:"composer"` // types : Organization Person

// FirstPerformance see : /firstPerformance
// The date and place the work was first performed.
FirstPerformance *Event `json:"firstPerformance"`

// IncludedComposition see : /includedComposition
// Smaller compositions included in this work (e.g. a movement in a symphony).
IncludedComposition *MusicComposition `json:"includedComposition"`

// IswcCode see : /iswcCode
// The International Standard Musical Work Code for the composition.
IswcCode string `json:"iswcCode"`

// Lyricist see : /lyricist
// The person who wrote the words.
Lyricist *Person `json:"lyricist"`

// Lyrics see : /lyrics
// The words in the song.
Lyrics *CreativeWork `json:"lyrics"`

// MusicArrangement see : /musicArrangement
// An arrangement derived from the composition.
MusicArrangement *MusicComposition `json:"musicArrangement"`

// MusicCompositionForm see : /musicCompositionForm
// The type of composition (e.g. overture, sonata, symphony, etc.).
MusicCompositionForm string `json:"musicCompositionForm"`

// MusicalKey see : /musicalKey
// The key, mode, or scale this composition uses.
MusicalKey string `json:"musicalKey"`

// RecordedAs see : /recordedAs
// An audio recording of the work. Inverse property: recordingOf (see: https://schema.org/recordingOf).
RecordedAs *MusicRecording `json:"recordedAs"`

}

func (v *MusicComposition) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicComposition"

	return json.Marshal(v)
}
