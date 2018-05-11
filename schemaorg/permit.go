package schemaorg

import "encoding/json"

// Permit see : https://schema.org/Permit
type Permit struct {

typeContext

Intangible

// IssuedBy see : /issuedBy
// The organization issuing the ticket or permit.
IssuedBy *Organization `json:"issuedBy"`

// IssuedThrough see : /issuedThrough
// The service through with the permit was granted.
IssuedThrough *Service `json:"issuedThrough"`

// PermitAudience see : /permitAudience
// The target audience for this permit.
PermitAudience *Audience `json:"permitAudience"`

// ValidFor see : /validFor
// The time validity of the permit.
ValidFor *Duration `json:"validFor"`

// ValidFrom see : /validFrom
// The date when the item becomes valid.
ValidFrom interface{} `json:"validFrom"`

// ValidIn see : /validIn
// The geographic area where the permit is valid.
ValidIn *AdministrativeArea `json:"validIn"`

// ValidUntil see : /validUntil
// The date when the item is no longer valid.
ValidUntil interface{} `json:"validUntil"`

}

func (v *Permit) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Permit"

	return json.Marshal(v)
}
