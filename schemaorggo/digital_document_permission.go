package schemaorggo

import "encoding/json"

// DigitalDocumentPermission see : https://schema.org/DigitalDocumentPermission
type DigitalDocumentPermission struct {
	Intangible

	typeContext

	// Grantee see : https://schema.org/grantee
	// The person, organization, contact point, or audience that has been granted this permission.
	// types : Audience ContactPoint Organization Person
	Grantee []interface{} `json:"grantee,omitempty"`

	// PermissionType see : https://schema.org/permissionType
	// The type of permission granted the person, organization, or audience.
	// types : DigitalDocumentPermissionType
	PermissionType []*DigitalDocumentPermissionType `json:"permissionType,omitempty"`
}

func (v DigitalDocumentPermission) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Grantee
		if len(v.Grantee) == 1 {
			value = v.Grantee[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["grantee"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PermissionType
		if len(v.PermissionType) == 1 {
			value = v.PermissionType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["permissionType"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v DigitalDocumentPermission) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DigitalDocumentPermission"

	return data, nil
}

func (v DigitalDocumentPermission) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
