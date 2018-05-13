package schemaorggo

import "encoding/json"

// ProgramMembership see : https://schema.org/ProgramMembership
type ProgramMembership struct {
	Intangible

	typeContext

	// HostingOrganization see : https://schema.org/hostingOrganization
	// The organization (airline, travelers&#39; club, etc.) the membership is made with.
	// types : Organization
	HostingOrganization []*Organization `json:"hostingOrganization,omitempty"`

	// Member see : https://schema.org/member
	// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	// types : Organization Person
	Member []interface{} `json:"member,omitempty"`

	// MembershipNumber see : https://schema.org/membershipNumber
	// A unique identifier for the membership.
	// types : Text
	MembershipNumber []string `json:"membershipNumber,omitempty"`

	// ProgramName see : https://schema.org/programName
	// The program providing the membership.
	// types : Text
	ProgramName []string `json:"programName,omitempty"`
}

func (v ProgramMembership) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.HostingOrganization
		if len(v.HostingOrganization) == 1 {
			value = v.HostingOrganization[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hostingOrganization"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Member
		if len(v.Member) == 1 {
			value = v.Member[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["member"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MembershipNumber
		if len(v.MembershipNumber) == 1 {
			value = v.MembershipNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["membershipNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProgramName
		if len(v.ProgramName) == 1 {
			value = v.ProgramName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["programName"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ProgramMembership) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ProgramMembership"

	return data, nil
}

func (v ProgramMembership) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
