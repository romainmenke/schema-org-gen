package schemaorggo

import "encoding/json"

// WebSite see : https://schema.org/WebSite
type WebSite struct {
	CreativeWork

	typeContext

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn string `json:"issn,omitempty"`
}

func (v WebSite) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WebSite"

	return json.Marshal(v)
}

func (v *WebSite) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WebSite"

	return json.Marshal(*v)
}
