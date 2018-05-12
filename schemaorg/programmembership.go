package schemaorg

import "encoding/json"

// ProgramMembership see : https://schema.org/ProgramMembership
type ProgramMembership struct {

Intangible

typeContext

// HostingOrganization see : https://schema.org/hostingOrganization
// The organization (airline, travelers' club, etc.) the membership is made with.
HostingOrganization *Organization `json:"hostingOrganization"`

// Member see : https://schema.org/member
// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
Member interface{} `json:"member"` // types : Organization Person

// MembershipNumber see : https://schema.org/membershipNumber
// A unique identifier for the membership.
MembershipNumber string `json:"membershipNumber"`

// ProgramName see : https://schema.org/programName
// The program providing the membership.
ProgramName string `json:"programName"`

}

func (v *ProgramMembership) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ProgramMembership"

	return json.Marshal(v)
}
