package schemaorg

import "encoding/json"

// OwnershipInfo see : https://schema.org/OwnershipInfo
type OwnershipInfo struct {

typeContext

StructuredValue

// AcquiredFrom see : /acquiredFrom
// The organization or person from which the product was acquired.
AcquiredFrom interface{} `json:"acquiredFrom"` // types : Organization Person

// OwnedFrom see : /ownedFrom
// The date and time of obtaining the product.
OwnedFrom interface{} `json:"ownedFrom"`

// OwnedThrough see : /ownedThrough
// The date and time of giving up ownership on the product.
OwnedThrough interface{} `json:"ownedThrough"`

// TypeOfGood see : /typeOfGood
// The product that this structured value is referring to.
TypeOfGood interface{} `json:"typeOfGood"` // types : Product Service

}

func (v *OwnershipInfo) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OwnershipInfo"

	return json.Marshal(v)
}
