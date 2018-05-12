package schemaorg

import "encoding/json"

// NoteDigitalDocument see : https://schema.org/NoteDigitalDocument
type NoteDigitalDocument struct {

DigitalDocument

typeContext

// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission"`

}

func (v *NoteDigitalDocument) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NoteDigitalDocument"

	return json.Marshal(v)
}
