package schemaorg

import "encoding/json"

// PresentationDigitalDocument see : https://schema.org/PresentationDigitalDocument
type PresentationDigitalDocument struct {

typeContext

DigitalDocument

// HasDigitalDocumentPermission see : /hasDigitalDocumentPermission
// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission"`

}

func (v *PresentationDigitalDocument) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PresentationDigitalDocument"

	return json.Marshal(v)
}
