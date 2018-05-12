package schemaorggo

import "encoding/json"

// DigitalDocumentPermissionType see : https://schema.org/DigitalDocumentPermissionType
type DigitalDocumentPermissionType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *DigitalDocumentPermissionType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DigitalDocumentPermissionType"

	return json.Marshal(v)
}
