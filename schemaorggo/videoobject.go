package schemaorggo

import "encoding/json"

// VideoObject see : https://schema.org/VideoObject
type VideoObject struct {
	MediaObject

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor,omitempty"`

	// Caption see : https://schema.org/caption
	// The caption for this object.
	Caption string `json:"caption,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy,omitempty"` // types : MusicGroup Person

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	Thumbnail *ImageObject `json:"thumbnail,omitempty"`

	// Transcript see : https://schema.org/transcript
	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	Transcript string `json:"transcript,omitempty"`

	// VideoFrameSize see : https://schema.org/videoFrameSize
	// The frame size of the video.
	VideoFrameSize string `json:"videoFrameSize,omitempty"`

	// VideoQuality see : https://schema.org/videoQuality
	// The quality of the video.
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