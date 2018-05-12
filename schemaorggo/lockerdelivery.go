package schemaorggo

import "encoding/json"

// LockerDelivery see : https://schema.org/LockerDelivery
type LockerDelivery struct {
	DeliveryMethod

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v LockerDelivery) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LockerDelivery"

	return json.Marshal(v)
}

func (v *LockerDelivery) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LockerDelivery"

	return json.Marshal(*v)
}
