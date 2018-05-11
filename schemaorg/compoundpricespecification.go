package schemaorg

import "encoding/json"

// CompoundPriceSpecification see : https://schema.org/CompoundPriceSpecification
type CompoundPriceSpecification struct {

typeContext

PriceSpecification

// PriceComponent see : https://schema.org/priceComponent
// This property links to all UnitPriceSpecification (see: https://schema.org/UnitPriceSpecification) nodes that apply in parallel for the CompoundPriceSpecification (see: https://schema.org/CompoundPriceSpecification) node.
PriceComponent *UnitPriceSpecification `json:"priceComponent"`

}

func (v *CompoundPriceSpecification) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CompoundPriceSpecification"

	return json.Marshal(v)
}
