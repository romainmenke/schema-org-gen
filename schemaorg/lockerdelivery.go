package schemaorg

import "encoding/json"

// LockerDelivery see : https://schema.org/LockerDelivery
type LockerDelivery struct {

DeliveryMethod

typeContext

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *LockerDelivery) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LockerDelivery"

	return json.Marshal(v)
}
