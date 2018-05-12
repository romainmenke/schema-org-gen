package schemaorggo

import "encoding/json"

// HowToTool see : https://schema.org/HowToTool
type HowToTool struct {
	HowToItem

	typeContext

	// RequiredQuantity see : https://schema.org/requiredQuantity
	// The required quantity of the item(s).
	// types : Number QuantitativeValue Text
	RequiredQuantity interface{} `json:"requiredQuantity,omitempty"`
}

func (v HowToTool) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToTool"

	return json.Marshal(v)
}

func (v *HowToTool) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HowToTool"

	return json.Marshal(*v)
}
