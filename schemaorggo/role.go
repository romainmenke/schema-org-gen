package schemaorggo

import "encoding/json"

// Role see : https://schema.org/Role
type Role struct {
	Intangible

	typeContext

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate []interface{} `json:"endDate,omitempty"`

	// RoleName see : https://schema.org/roleName
	// A role played, performed or filled by a person or organization. For example, the team of creators for a comic book might fill the roles named &#39;inker&#39;, &#39;penciller&#39;, and &#39;letterer&#39;; or an athlete in a SportsTeam might play in the position named &#39;Quarterback&#39;. Supersedes namedPosition (see: https://schema.org/namedPosition).
	// types : Text URL
	RoleName []string `json:"roleName,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate []interface{} `json:"startDate,omitempty"`
}

func (v Role) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.EndDate
		if len(v.EndDate) == 1 {
			value = v.EndDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RoleName
		if len(v.RoleName) == 1 {
			value = v.RoleName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["roleName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StartDate
		if len(v.StartDate) == 1 {
			value = v.StartDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["startDate"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Role) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Role"

	return data, nil
}

func (v Role) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
