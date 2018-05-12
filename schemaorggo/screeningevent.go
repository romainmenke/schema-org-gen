package schemaorggo

import "encoding/json"

// ScreeningEvent see : https://schema.org/ScreeningEvent
type ScreeningEvent struct {
	Event

	typeContext

	// SubtitleLanguage see : https://schema.org/subtitleLanguage
	// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	// types : Language Text
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat string `json:"videoFormat,omitempty"`

	// WorkPresented see : https://schema.org/workPresented
	// The movie presented during this event.
	// types : Movie
	WorkPresented *Movie `json:"workPresented,omitempty"`
}

func (v ScreeningEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ScreeningEvent"

	return json.Marshal(v)
}

func (v *ScreeningEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ScreeningEvent"

	return json.Marshal(*v)
}
