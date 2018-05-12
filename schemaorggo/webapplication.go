package schemaorggo

import "encoding/json"

// WebApplication see : https://schema.org/WebApplication
type WebApplication struct {
	SoftwareApplication

	typeContext

	// BrowserRequirements see : https://schema.org/browserRequirements
	// Specifies browser requirements in human-readable text. For example, 'requires HTML5 support'.
	BrowserRequirements string `json:"browserRequirements"`
}

func (v *WebApplication) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WebApplication"

	return json.Marshal(v)
}
