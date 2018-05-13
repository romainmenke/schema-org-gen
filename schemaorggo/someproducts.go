package schemaorggo

import "encoding/json"

// SomeProducts see : https://schema.org/SomeProducts
type SomeProducts struct {
	Product

	typeContext

	// InventoryLevel see : https://schema.org/inventoryLevel
	// The current approximate inventory level for the item or items.
	// types : QuantitativeValue
	InventoryLevel []*QuantitativeValue `json:"inventoryLevel,omitempty"`
}

func (v SomeProducts) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Product.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.InventoryLevel
		if len(v.InventoryLevel) == 1 {
			value = v.InventoryLevel[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inventoryLevel"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SomeProducts) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SomeProducts"

	return data, nil
}

func (v SomeProducts) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
