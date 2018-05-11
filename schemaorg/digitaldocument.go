package schemaorg

import "encoding/json"

// DigitalDocument see : https://schema.org/DigitalDocument
type DigitalDocument struct {

typeContext

CreativeWork

// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission"`

}

func (v *DigitalDocument) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DigitalDocument"

	return json.Marshal(v)
}
