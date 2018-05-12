package schemaorggo

import "encoding/json"

// WriteAction see : https://schema.org/WriteAction
type WriteAction struct {
	CreateAction

	typeContext

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage interface{} `json:"inLanguage,omitempty"`
}

func (v WriteAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WriteAction"

	return json.Marshal(v)
}

func (v *WriteAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WriteAction"

	return json.Marshal(*v)
}
