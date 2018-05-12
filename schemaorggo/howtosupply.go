package schemaorggo

import "encoding/json"

// HowToSupply see : https://schema.org/HowToSupply
type HowToSupply struct {
	HowToItem

	typeContext

	// EstimatedCost see : https://schema.org/estimatedCost
	// The estimated cost of the supply or supplies consumed when performing instructions.
	EstimatedCost interface{} `json:"estimatedCost,omitempty"` // types : MonetaryAmount Text

}

func (v HowToSupply) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToSupply"

	return json.Marshal(v)
}

func (v *HowToSupply) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HowToSupply"

	return json.Marshal(*v)
}
