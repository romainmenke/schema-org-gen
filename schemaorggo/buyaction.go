package schemaorggo

import "encoding/json"

// BuyAction see : https://schema.org/BuyAction
type BuyAction struct {
	TradeAction

	typeContext

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller interface{} `json:"seller,omitempty"`
}

func (v BuyAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BuyAction"

	return json.Marshal(v)
}

func (v *BuyAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BuyAction"

	return json.Marshal(*v)
}
