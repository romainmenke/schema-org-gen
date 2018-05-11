package schemaorg

import "encoding/json"

// BuyAction see : https://schema.org/BuyAction
type BuyAction struct {

typeContext

TradeAction

// Seller see : /seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
Seller interface{} `json:"seller"` // types : Organization Person

}

func (v *BuyAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BuyAction"

	return json.Marshal(v)
}
