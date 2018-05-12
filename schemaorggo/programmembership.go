package schemaorggo

import "encoding/json"

// ProgramMembership see : https://schema.org/ProgramMembership
type ProgramMembership struct {
	Intangible

	typeContext

	// HostingOrganization see : https://schema.org/hostingOrganization
	// The organization (airline, travelers&#39; club, etc.) the membership is made with.
	// types : Organization
	HostingOrganization *Organization `json:"hostingOrganization,omitempty"`

	// Member see : https://schema.org/member
	// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	// types : Organization Person
	Member interface{} `json:"member,omitempty"`

	// MembershipNumber see : https://schema.org/membershipNumber
	// A unique identifier for the membership.
	// types : Text
	MembershipNumber string `json:"membershipNumber,omitempty"`

	// ProgramName see : https://schema.org/programName
	// The program providing the membership.
	// types : Text
	ProgramName string `json:"programName,omitempty"`
}

func (v ProgramMembership) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ProgramMembership"

	return json.Marshal(v)
}

func (v *ProgramMembership) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ProgramMembership"

	return json.Marshal(*v)
}
