package schemaorg

import "encoding/json"

// HowToTool see : https://schema.org/HowToTool
type HowToTool struct {

typeContext

HowToItem

// RequiredQuantity see : https://schema.org/requiredQuantity
// The required quantity of the item(s).
RequiredQuantity interface{} `json:"requiredQuantity"` // types : Number QuantitativeValue Text

}

func (v *HowToTool) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToTool"

	return json.Marshal(v)
}
