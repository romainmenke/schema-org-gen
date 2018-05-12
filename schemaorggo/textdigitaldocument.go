package schemaorggo

import "encoding/json"

// TextDigitalDocument see : https://schema.org/TextDigitalDocument
type TextDigitalDocument struct {
	DigitalDocument

	typeContext

	// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
	// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty"`
}

func (v TextDigitalDocument) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TextDigitalDocument"

	return json.Marshal(v)
}

func (v *TextDigitalDocument) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TextDigitalDocument"

	return json.Marshal(*v)
}
