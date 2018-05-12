package schemaorggo

import "encoding/json"

// NoteDigitalDocument see : https://schema.org/NoteDigitalDocument
type NoteDigitalDocument struct {
	DigitalDocument

	typeContext

	// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
	// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to &quot;public&quot;.
	// types : DigitalDocumentPermission
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty"`
}

func (v NoteDigitalDocument) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NoteDigitalDocument"

	return json.Marshal(v)
}

func (v *NoteDigitalDocument) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "NoteDigitalDocument"

	return json.Marshal(*v)
}
