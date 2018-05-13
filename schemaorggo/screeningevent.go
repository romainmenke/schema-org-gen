package schemaorggo

import "encoding/json"

// ScreeningEvent see : https://schema.org/ScreeningEvent
type ScreeningEvent struct {
	Event

	typeContext

	// SubtitleLanguage see : https://schema.org/subtitleLanguage
	// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	// types : Language Text
	SubtitleLanguage []interface{} `json:"subtitleLanguage,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat []string `json:"videoFormat,omitempty"`

	// WorkPresented see : https://schema.org/workPresented
	// The movie presented during this event.
	// types : Movie
	WorkPresented []*Movie `json:"workPresented,omitempty"`
}

func (v ScreeningEvent) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Event.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.SubtitleLanguage
		if len(v.SubtitleLanguage) == 1 {
			value = v.SubtitleLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subtitleLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VideoFormat
		if len(v.VideoFormat) == 1 {
			value = v.VideoFormat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["videoFormat"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkPresented
		if len(v.WorkPresented) == 1 {
			value = v.WorkPresented[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workPresented"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ScreeningEvent) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ScreeningEvent"

	return data, nil
}

func (v ScreeningEvent) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
