package schemaorg

import "encoding/json"

// HowToItem see : https://schema.org/HowToItem
type HowToItem struct {

typeContext

ListItem

// RequiredQuantity see : https://schema.org/requiredQuantity
// The required quantity of the item(s).
RequiredQuantity interface{} `json:"requiredQuantity"` // types : Number QuantitativeValue Text

}

func (v *HowToItem) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToItem"

	return json.Marshal(v)
}
