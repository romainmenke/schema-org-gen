package schemaorggo

import "encoding/json"

// VideoObject see : https://schema.org/VideoObject
type VideoObject struct {
	MediaObject

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// Caption see : https://schema.org/caption
	// The caption for this object.
	// types : Text
	Caption []string `json:"caption,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy []interface{} `json:"musicBy,omitempty"`

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	// types : ImageObject
	Thumbnail []*ImageObject `json:"thumbnail,omitempty"`

	// Transcript see : https://schema.org/transcript
	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	// types : Text
	Transcript []string `json:"transcript,omitempty"`

	// VideoFrameSize see : https://schema.org/videoFrameSize
	// The frame size of the video.
	// types : Text
	VideoFrameSize []string `json:"videoFrameSize,omitempty"`

	// VideoQuality see : https://schema.org/videoQuality
	// The quality of the video.
	// types : Text
	VideoQuality []string `json:"videoQuality,omitempty"`
}

func (v VideoObject) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.MediaObject.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Actor
		if len(v.Actor) == 1 {
			value = v.Actor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Caption
		if len(v.Caption) == 1 {
			value = v.Caption[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["caption"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Director
		if len(v.Director) == 1 {
			value = v.Director[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["director"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicBy
		if len(v.MusicBy) == 1 {
			value = v.MusicBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Thumbnail
		if len(v.Thumbnail) == 1 {
			value = v.Thumbnail[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["thumbnail"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Transcript
		if len(v.Transcript) == 1 {
			value = v.Transcript[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["transcript"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VideoFrameSize
		if len(v.VideoFrameSize) == 1 {
			value = v.VideoFrameSize[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["videoFrameSize"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VideoQuality
		if len(v.VideoQuality) == 1 {
			value = v.VideoQuality[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["videoQuality"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v VideoObject) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "VideoObject"

	return data, nil
}

func (v VideoObject) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
