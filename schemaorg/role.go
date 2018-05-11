package schemaorg

import "encoding/json"

// Role see : https://schema.org/Role
type Role struct {

typeContext

Intangible

// EndDate see : https://schema.org/endDate
// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
EndDate interface{} `json:"endDate"` // types : Date DateTime

// RoleName see : https://schema.org/roleName
// A role played, performed or filled by a person or organization. For example, the team of creators for a comic book might fill the roles named 'inker', 'penciller', and 'letterer'; or an athlete in a SportsTeam might play in the position named 'Quarterback'. Supersedes namedPosition (see: https://schema.org/namedPosition).
RoleName interface{} `json:"roleName"` // types : Text URL

// StartDate see : https://schema.org/startDate
// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
StartDate interface{} `json:"startDate"` // types : Date DateTime

}

func (v *Role) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Role"

	return json.Marshal(v)
}
