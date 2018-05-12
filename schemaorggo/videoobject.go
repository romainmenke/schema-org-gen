package schemaorggo

import "encoding/json"

// VideoObject see : https://schema.org/VideoObject
type VideoObject struct {
	MediaObject

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// Caption see : https://schema.org/caption
	// The caption for this object.
	// types : Text
	Caption string `json:"caption,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy interface{} `json:"musicBy,omitempty"`

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	// types : ImageObject
	Thumbnail *ImageObject `json:"thumbnail,omitempty"`

	// Transcript see : https://schema.org/transcript
	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	// types : Text
	Transcript string `json:"transcript,omitempty"`

	// VideoFrameSize see : https://schema.org/videoFrameSize
	// The frame size of the video.
	// types : Text
	VideoFrameSize string `json:"videoFrameSize,omitempty"`

	// VideoQuality see : https://schema.org/videoQuality
	// The quality of the video.
	// types : Text
	VideoQuality string `json:"videoQuality,omitempty"`
}

func (v VideoObject) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoObject"

	return json.Marshal(v)
}

func (v *VideoObject) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "VideoObject"

	return json.Marshal(*v)
}
