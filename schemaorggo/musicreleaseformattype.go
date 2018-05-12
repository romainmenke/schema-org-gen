package schemaorggo

import "encoding/json"

// MusicReleaseFormatType see : https://schema.org/MusicReleaseFormatType
type MusicReleaseFormatType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v MusicReleaseFormatType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicReleaseFormatType"

	return json.Marshal(v)
}

func (v *MusicReleaseFormatType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MusicReleaseFormatType"

	return json.Marshal(*v)
}
