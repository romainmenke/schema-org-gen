package schemaorggo

import "encoding/json"

// OwnershipInfo see : https://schema.org/OwnershipInfo
type OwnershipInfo struct {
	StructuredValue

	typeContext

	// AcquiredFrom see : https://schema.org/acquiredFrom
	// The organization or person from which the product was acquired.
	AcquiredFrom interface{} `json:"acquiredFrom"` // types : Organization Person

	// OwnedFrom see : https://schema.org/ownedFrom
	// The date and time of obtaining the product.
	OwnedFrom DateTime `json:"ownedFrom"`

	// OwnedThrough see : https://schema.org/ownedThrough
	// The date and time of giving up ownership on the product.
	OwnedThrough DateTime `json:"ownedThrough"`

	// TypeOfGood see : https://schema.org/typeOfGood
	// The product that this structured value is referring to.
	TypeOfGood interface{} `json:"typeOfGood"` // types : Product Service

}

func (v *OwnershipInfo) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OwnershipInfo"

	return json.Marshal(v)
}
