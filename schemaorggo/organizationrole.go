package schemaorggo

import "encoding/json"

// OrganizationRole see : https://schema.org/OrganizationRole
type OrganizationRole struct {
	Role

	typeContext

	// NumberedPosition see : https://schema.org/numberedPosition
	// A number associated with a role in an organization, for example, the number on an athlete's jersey.
	NumberedPosition float64 `json:"numberedPosition,omitempty"` // types : Number

}

func (v OrganizationRole) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrganizationRole"

	return json.Marshal(v)
}

func (v *OrganizationRole) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OrganizationRole"

	return json.Marshal(*v)
}
