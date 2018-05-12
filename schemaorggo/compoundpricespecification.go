package schemaorggo

import "encoding/json"

// CompoundPriceSpecification see : https://schema.org/CompoundPriceSpecification
type CompoundPriceSpecification struct {
	PriceSpecification

	typeContext

	// PriceComponent see : https://schema.org/priceComponent
	// This property links to all UnitPriceSpecification (see: https://schema.org/UnitPriceSpecification) nodes that apply in parallel for the CompoundPriceSpecification (see: https://schema.org/CompoundPriceSpecification) node.
	PriceComponent *UnitPriceSpecification `json:"priceComponent,omitempty"` // types : UnitPriceSpecification

}

func (v CompoundPriceSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CompoundPriceSpecification"

	return json.Marshal(v)
}

func (v *CompoundPriceSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CompoundPriceSpecification"

	return json.Marshal(*v)
}
