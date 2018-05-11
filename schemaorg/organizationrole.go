package schemaorg

import "encoding/json"

// OrganizationRole see : https://schema.org/OrganizationRole
type OrganizationRole struct {

typeContext

Role

// NumberedPosition see : /numberedPosition
// A number associated with a role in an organization, for example, the number on an athlete's jersey.
NumberedPosition float64 `json:"numberedPosition"`

}

func (v *OrganizationRole) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrganizationRole"

	return json.Marshal(v)
}
