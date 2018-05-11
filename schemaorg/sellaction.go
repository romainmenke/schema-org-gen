package schemaorg

import "encoding/json"

// SellAction see : https://schema.org/SellAction
type SellAction struct {

typeContext

TradeAction

// Buyer see : /buyer
// A sub property of participant. The participant/person/organization that bought the object.
Buyer *Person `json:"buyer"`

}

func (v *SellAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SellAction"

	return json.Marshal(v)
}
