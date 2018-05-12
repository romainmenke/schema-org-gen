package schemaorggo

import "encoding/json"

// DigitalDocumentPermissionType see : https://schema.org/DigitalDocumentPermissionType
type DigitalDocumentPermissionType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v DigitalDocumentPermissionType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DigitalDocumentPermissionType"

	return json.Marshal(v)
}

func (v *DigitalDocumentPermissionType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DigitalDocumentPermissionType"

	return json.Marshal(*v)
}
