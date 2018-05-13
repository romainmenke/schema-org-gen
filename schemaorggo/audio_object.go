package schemaorggo

import "encoding/json"

// AudioObject see : https://schema.org/AudioObject
type AudioObject struct {
	MediaObject

	typeContext

	// Transcript see : https://schema.org/transcript
	// If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	// types : Text
	Transcript []string `json:"transcript,omitempty"`
}

func (v AudioObject) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.MediaObject.intoMap(intop)

	into := *intop

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

	*intop = into

	return nil
}

func (v AudioObject) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AudioObject"

	return data, nil
}

func (v AudioObject) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
