package schemaorggo

import "encoding/json"

// CompoundPriceSpecification see : https://schema.org/CompoundPriceSpecification
type CompoundPriceSpecification struct {
	PriceSpecification

	typeContext

	// PriceComponent see : https://schema.org/priceComponent
	// This property links to all UnitPriceSpecification (see: https://schema.org/UnitPriceSpecification) nodes that apply in parallel for the CompoundPriceSpecification (see: https://schema.org/CompoundPriceSpecification) node.
	// types : UnitPriceSpecification
	PriceComponent []*UnitPriceSpecification `json:"priceComponent,omitempty"`
}

func (v CompoundPriceSpecification) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PriceSpecification.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.PriceComponent
		if len(v.PriceComponent) == 1 {
			value = v.PriceComponent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceComponent"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v CompoundPriceSpecification) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CompoundPriceSpecification"

	return data, nil
}

func (v CompoundPriceSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
