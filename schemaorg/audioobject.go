package schemaorg

import "encoding/json"

// AudioObject see : https://schema.org/AudioObject
type AudioObject struct {

typeContext

MediaObject

// Transcript see : /transcript
// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
Transcript string `json:"transcript"`

}

func (v *AudioObject) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AudioObject"

	return json.Marshal(v)
}
