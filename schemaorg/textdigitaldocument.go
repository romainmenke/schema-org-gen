package schemaorg

import "encoding/json"

// TextDigitalDocument see : https://schema.org/TextDigitalDocument
type TextDigitalDocument struct {

typeContext

DigitalDocument

// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission"`

}

func (v *TextDigitalDocument) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TextDigitalDocument"

	return json.Marshal(v)
}
