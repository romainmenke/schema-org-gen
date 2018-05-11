package schemaorg

import "encoding/json"

// SomeProducts see : https://schema.org/SomeProducts
type SomeProducts struct {

typeContext

Product

// InventoryLevel see : https://schema.org/inventoryLevel
// The current approximate inventory level for the item or items.
InventoryLevel *QuantitativeValue `json:"inventoryLevel"`

}

func (v *SomeProducts) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SomeProducts"

	return json.Marshal(v)
}
