package schemaorggo

import "encoding/json"

// OwnershipInfo see : https://schema.org/OwnershipInfo
type OwnershipInfo struct {
	StructuredValue

	typeContext

	// AcquiredFrom see : https://schema.org/acquiredFrom
	// The organization or person from which the product was acquired.
	AcquiredFrom interface{} `json:"acquiredFrom,omitempty"` // types : Organization Person

	// OwnedFrom see : https://schema.org/ownedFrom
	// The date and time of obtaining the product.
	OwnedFrom DateTime `json:"ownedFrom,omitempty"` // types : DateTime

	// OwnedThrough see : https://schema.org/ownedThrough
	// The date and time of giving up ownership on the product.
	OwnedThrough DateTime `json:"ownedThrough,omitempty"` // types : DateTime

	// TypeOfGood see : https://schema.org/typeOfGood
	// The product that this structured value is referring to.
	TypeOfGood interface{} `json:"typeOfGood,omitempty"` // types : Product Service

}

func (v OwnershipInfo) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OwnershipInfo"

	return json.Marshal(v)
}

func (v *OwnershipInfo) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OwnershipInfo"

	return json.Marshal(*v)
}
