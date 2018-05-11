package schemaorg

import "encoding/json"

// ScreeningEvent see : https://schema.org/ScreeningEvent
type ScreeningEvent struct {

typeContext

Event

// SubtitleLanguage see : https://schema.org/subtitleLanguage
// Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
SubtitleLanguage interface{} `json:"subtitleLanguage"` // types : Language Text

// VideoFormat see : https://schema.org/videoFormat
// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
VideoFormat string `json:"videoFormat"`

// WorkPresented see : https://schema.org/workPresented
// The movie presented during this event.
WorkPresented *Movie `json:"workPresented"`

}

func (v *ScreeningEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ScreeningEvent"

	return json.Marshal(v)
}
