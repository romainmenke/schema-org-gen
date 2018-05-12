package schemaorggo

import "encoding/json"

// Permit see : https://schema.org/Permit
type Permit struct {
	Intangible

	typeContext

	// IssuedBy see : https://schema.org/issuedBy
	// The organization issuing the ticket or permit.
	IssuedBy *Organization `json:"issuedBy,omitempty"` // types : Organization

	// IssuedThrough see : https://schema.org/issuedThrough
	// The service through with the permit was granted.
	IssuedThrough *Service `json:"issuedThrough,omitempty"` // types : Service

	// PermitAudience see : https://schema.org/permitAudience
	// The target audience for this permit.
	PermitAudience *Audience `json:"permitAudience,omitempty"` // types : Audience

	// ValidFor see : https://schema.org/validFor
	// The time validity of the permit.
	ValidFor *Duration `json:"validFor,omitempty"` // types : Duration

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	ValidFrom DateTime `json:"validFrom,omitempty"` // types : DateTime

	// ValidIn see : https://schema.org/validIn
	// The geographic area where the permit is valid.
	ValidIn *AdministrativeArea `json:"validIn,omitempty"` // types : AdministrativeArea

	// ValidUntil see : https://schema.org/validUntil
	// The date when the item is no longer valid.
	ValidUntil Date `json:"validUntil,omitempty"` // types : Date

}

func (v Permit) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Permit"

	return json.Marshal(v)
}

func (v *Permit) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Permit"

	return json.Marshal(*v)
}
