package schemaorg

import "encoding/json"

// Permit see : https://schema.org/Permit
type Permit struct {

Intangible

typeContext

// IssuedBy see : https://schema.org/issuedBy
// The organization issuing the ticket or permit.
IssuedBy *Organization `json:"issuedBy"`

// IssuedThrough see : https://schema.org/issuedThrough
// The service through with the permit was granted.
IssuedThrough *Service `json:"issuedThrough"`

// PermitAudience see : https://schema.org/permitAudience
// The target audience for this permit.
PermitAudience *Audience `json:"permitAudience"`

// ValidFor see : https://schema.org/validFor
// The time validity of the permit.
ValidFor *Duration `json:"validFor"`

// ValidFrom see : https://schema.org/validFrom
// The date when the item becomes valid.
ValidFrom interface{} `json:"validFrom"`

// ValidIn see : https://schema.org/validIn
// The geographic area where the permit is valid.
ValidIn *AdministrativeArea `json:"validIn"`

// ValidUntil see : https://schema.org/validUntil
// The date when the item is no longer valid.
ValidUntil interface{} `json:"validUntil"`

}

func (v *Permit) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Permit"

	return json.Marshal(v)
}
