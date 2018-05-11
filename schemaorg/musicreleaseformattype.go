package schemaorg

import "encoding/json"

// MusicReleaseFormatType see : https://schema.org/MusicReleaseFormatType
type MusicReleaseFormatType struct {

typeContext

Enumeration

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *MusicReleaseFormatType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicReleaseFormatType"

	return json.Marshal(v)
}
