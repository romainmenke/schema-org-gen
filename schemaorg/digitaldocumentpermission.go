package schemaorg

import "encoding/json"

// DigitalDocumentPermission see : https://schema.org/DigitalDocumentPermission
type DigitalDocumentPermission struct {

typeContext

Intangible

// Grantee see : https://schema.org/grantee
// The person, organization, contact point, or audience that has been granted this permission.
Grantee interface{} `json:"grantee"` // types : Audience ContactPoint Organization Person

// PermissionType see : https://schema.org/permissionType
// The type of permission granted the person, organization, or audience.
PermissionType *DigitalDocumentPermissionType `json:"permissionType"`

}

func (v *DigitalDocumentPermission) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DigitalDocumentPermission"

	return json.Marshal(v)
}
