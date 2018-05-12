package schemaorggo

import "encoding/json"

// GovernmentPermit see : https://schema.org/GovernmentPermit
type GovernmentPermit struct {
	Permit

	typeContext

	// IssuedBy see : https://schema.org/issuedBy
	// The organization issuing the ticket or permit.
	// types : Organization
	IssuedBy *Organization `json:"issuedBy,omitempty"`

	// IssuedThrough see : https://schema.org/issuedThrough
	// The service through with the permit was granted.
	// types : Service
	IssuedThrough *Service `json:"issuedThrough,omitempty"`

	// PermitAudience see : https://schema.org/permitAudience
	// The target audience for this permit.
	// types : Audience
	PermitAudience *Audience `json:"permitAudience,omitempty"`

	// ValidFor see : https://schema.org/validFor
	// The time validity of the permit.
	// types : Duration
	ValidFor *Duration `json:"validFor,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom DateTime `json:"validFrom,omitempty"`

	// ValidIn see : https://schema.org/validIn
	// The geographic area where the permit is valid.
	// types : AdministrativeArea
	ValidIn *AdministrativeArea `json:"validIn,omitempty"`

	// ValidUntil see : https://schema.org/validUntil
	// The date when the item is no longer valid.
	// types : Date
	ValidUntil Date `json:"validUntil,omitempty"`
}

func (v GovernmentPermit) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GovernmentPermit"

	return json.Marshal(v)
}

func (v *GovernmentPermit) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GovernmentPermit"

	return json.Marshal(*v)
}
