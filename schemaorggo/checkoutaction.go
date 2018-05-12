package schemaorggo

import "encoding/json"

// CheckOutAction see : https://schema.org/CheckOutAction
type CheckOutAction struct {
	CommunicateAction

	typeContext

	// About see : https://schema.org/about
	// The subject matter of the content.
	About *Thing `json:"about,omitempty"` // types : Thing

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	InLanguage interface{} `json:"inLanguage,omitempty"` // types : Language Text

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	Recipient interface{} `json:"recipient,omitempty"` // types : Audience ContactPoint Organization Person

}

func (v CheckOutAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CheckOutAction"

	return json.Marshal(v)
}

func (v *CheckOutAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CheckOutAction"

	return json.Marshal(*v)
}
