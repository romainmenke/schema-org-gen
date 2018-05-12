package schemaorggo

import "encoding/json"

// ScreeningEvent see : https://schema.org/ScreeningEvent
type ScreeningEvent struct {
	Event

	typeContext

	// SubtitleLanguage see : https://schema.org/subtitleLanguage
	// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty"` // types : Language Text

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	VideoFormat string `json:"videoFormat,omitempty"` // types : Text

	// WorkPresented see : https://schema.org/workPresented
	// The movie presented during this event.
	WorkPresented *Movie `json:"workPresented,omitempty"` // types : Movie

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
