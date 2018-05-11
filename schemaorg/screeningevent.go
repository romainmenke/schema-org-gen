package schemaorg

import "encoding/json"

// ScreeningEvent see : https://schema.org/ScreeningEvent
type ScreeningEvent struct {

typeContext

Event

// SubtitleLanguage see : /subtitleLanguage
// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
SubtitleLanguage interface{} `json:"subtitleLanguage"` // types : Language Text

// VideoFormat see : /videoFormat
// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
VideoFormat string `json:"videoFormat"`

// WorkPresented see : /workPresented
// The movie presented during this event.
WorkPresented *Movie `json:"workPresented"`

}

func (v *ScreeningEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ScreeningEvent"

	return json.Marshal(v)
}
