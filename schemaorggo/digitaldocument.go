package schemaorggo

import "encoding/json"

// DigitalDocument see : https://schema.org/DigitalDocument
type DigitalDocument struct {
	CreativeWork

	typeContext

	// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
	// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty"`
}

func (v DigitalDocument) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DigitalDocument"

	return json.Marshal(v)
}

func (v *DigitalDocument) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DigitalDocument"

	return json.Marshal(*v)
}
