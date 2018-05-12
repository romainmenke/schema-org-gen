package schemaorggo

import "encoding/json"

// DigitalDocumentPermission see : https://schema.org/DigitalDocumentPermission
type DigitalDocumentPermission struct {
	Intangible

	typeContext

	// Grantee see : https://schema.org/grantee
	// The person, organization, contact point, or audience that has been granted this permission.
	Grantee interface{} `json:"grantee,omitempty"` // types : Audience ContactPoint Organization Person

	// PermissionType see : https://schema.org/permissionType
	// The type of permission granted the person, organization, or audience.
	PermissionType *DigitalDocumentPermissionType `json:"permissionType,omitempty"`
}

func (v DigitalDocumentPermission) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DigitalDocumentPermission"

	return json.Marshal(v)
}

func (v *DigitalDocumentPermission) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DigitalDocumentPermission"

	return json.Marshal(*v)
}
