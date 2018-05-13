package schemaorggo

import "encoding/json"

// GovernmentPermit see : https://schema.org/GovernmentPermit
type GovernmentPermit struct {
	Permit

	typeContext

	// IssuedBy see : https://schema.org/issuedBy
	// The organization issuing the ticket or permit.
	// types : Organization
	IssuedBy []*Organization `json:"issuedBy,omitempty"`

	// IssuedThrough see : https://schema.org/issuedThrough
	// The service through with the permit was granted.
	// types : Service
	IssuedThrough []*Service `json:"issuedThrough,omitempty"`

	// PermitAudience see : https://schema.org/permitAudience
	// The target audience for this permit.
	// types : Audience
	PermitAudience []*Audience `json:"permitAudience,omitempty"`

	// ValidFor see : https://schema.org/validFor
	// The time validity of the permit.
	// types : Duration
	ValidFor []*Duration `json:"validFor,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom []DateTime `json:"validFrom,omitempty"`

	// ValidIn see : https://schema.org/validIn
	// The geographic area where the permit is valid.
	// types : AdministrativeArea
	ValidIn []*AdministrativeArea `json:"validIn,omitempty"`

	// ValidUntil see : https://schema.org/validUntil
	// The date when the item is no longer valid.
	// types : Date
	ValidUntil []Date `json:"validUntil,omitempty"`
}

func (v GovernmentPermit) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Permit.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.IssuedBy
		if len(v.IssuedBy) == 1 {
			value = v.IssuedBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["issuedBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IssuedThrough
		if len(v.IssuedThrough) == 1 {
			value = v.IssuedThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["issuedThrough"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PermitAudience
		if len(v.PermitAudience) == 1 {
			value = v.PermitAudience[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["permitAudience"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidFor
		if len(v.ValidFor) == 1 {
			value = v.ValidFor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validFor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidFrom
		if len(v.ValidFrom) == 1 {
			value = v.ValidFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidIn
		if len(v.ValidIn) == 1 {
			value = v.ValidIn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validIn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidUntil
		if len(v.ValidUntil) == 1 {
			value = v.ValidUntil[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validUntil"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v GovernmentPermit) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "GovernmentPermit"

	return data, nil
}

func (v GovernmentPermit) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
