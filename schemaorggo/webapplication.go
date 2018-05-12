package schemaorggo

import "encoding/json"

// WebApplication see : https://schema.org/WebApplication
type WebApplication struct {
	SoftwareApplication

	typeContext

	// BrowserRequirements see : https://schema.org/browserRequirements
	// Specifies browser requirements in human-readable text. For example, &#39;requires HTML5 support&#39;.
	// types : Text
	BrowserRequirements string `json:"browserRequirements,omitempty"`
}

func (v WebApplication) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WebApplication"

	return json.Marshal(v)
}

func (v *WebApplication) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WebApplication"

	return json.Marshal(*v)
}
