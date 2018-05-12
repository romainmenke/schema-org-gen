package schemaorggo

import "encoding/json"

// Role see : https://schema.org/Role
type Role struct {
	Intangible

	typeContext

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate interface{} `json:"endDate,omitempty"`

	// RoleName see : https://schema.org/roleName
	// A role played, performed or filled by a person or organization. For example, the team of creators for a comic book might fill the roles named &#39;inker&#39;, &#39;penciller&#39;, and &#39;letterer&#39;; or an athlete in a SportsTeam might play in the position named &#39;Quarterback&#39;. Supersedes namedPosition (see: https://schema.org/namedPosition).
	// types : Text URL
	RoleName string `json:"roleName,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate interface{} `json:"startDate,omitempty"`
}

func (v Role) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Role"

	return json.Marshal(v)
}

func (v *Role) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Role"

	return json.Marshal(*v)
}
