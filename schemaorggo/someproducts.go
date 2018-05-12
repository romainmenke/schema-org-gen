package schemaorggo

import "encoding/json"

// SomeProducts see : https://schema.org/SomeProducts
type SomeProducts struct {
	Product

	typeContext

	// InventoryLevel see : https://schema.org/inventoryLevel
	// The current approximate inventory level for the item or items.
	InventoryLevel *QuantitativeValue `json:"inventoryLevel,omitempty"`
}

func (v SomeProducts) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SomeProducts"

	return json.Marshal(v)
}

func (v *SomeProducts) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SomeProducts"

	return json.Marshal(*v)
}
