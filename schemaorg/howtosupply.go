package schemaorg

import "encoding/json"

// HowToSupply see : https://schema.org/HowToSupply
type HowToSupply struct {

typeContext

HowToItem

// EstimatedCost see : /estimatedCost
// The estimated cost of the supply or supplies consumed when performing instructions.
EstimatedCost interface{} `json:"estimatedCost"` // types : MonetaryAmount Text

}

func (v *HowToSupply) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToSupply"

	return json.Marshal(v)
}
