package schemaorggo

import "encoding/json"

// ShareAction see : https://schema.org/ShareAction
type ShareAction struct {
	CommunicateAction

	typeContext

	// About see : https://schema.org/about
	// The subject matter of the content.
	// types : Thing
	About *Thing `json:"about,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage interface{} `json:"inLanguage,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`
}

func (v ShareAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ShareAction"

	return json.Marshal(v)
}

func (v *ShareAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ShareAction"

	return json.Marshal(*v)
}
