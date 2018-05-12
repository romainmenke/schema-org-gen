package schemaorggo

import "encoding/json"

// HowToItem see : https://schema.org/HowToItem
type HowToItem struct {
	ListItem

	typeContext

	// RequiredQuantity see : https://schema.org/requiredQuantity
	// The required quantity of the item(s).
	RequiredQuantity interface{} `json:"requiredQuantity,omitempty"` // types : Number QuantitativeValue Text

}

func (v HowToItem) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToItem"

	return json.Marshal(v)
}

func (v *HowToItem) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HowToItem"

	return json.Marshal(*v)
}
