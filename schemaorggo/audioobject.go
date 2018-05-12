package schemaorggo

import "encoding/json"

// AudioObject see : https://schema.org/AudioObject
type AudioObject struct {
	MediaObject

	typeContext

	// Transcript see : https://schema.org/transcript
	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	// types : Text
	Transcript string `json:"transcript,omitempty"`
}

func (v AudioObject) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AudioObject"

	return json.Marshal(v)
}

func (v *AudioObject) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AudioObject"

	return json.Marshal(*v)
}
