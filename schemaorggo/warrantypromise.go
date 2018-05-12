package schemaorggo

import "encoding/json"

// WarrantyPromise see : https://schema.org/WarrantyPromise
type WarrantyPromise struct {
	StructuredValue

	typeContext

	// DurationOfWarranty see : https://schema.org/durationOfWarranty
	// The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days.
	// types : QuantitativeValue
	DurationOfWarranty *QuantitativeValue `json:"durationOfWarranty,omitempty"`

	// WarrantyScope see : https://schema.org/warrantyScope
	// The scope of the warranty promise.
	// types : WarrantyScope
	WarrantyScope *WarrantyScope `json:"warrantyScope,omitempty"`
}

func (v WarrantyPromise) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WarrantyPromise"

	return json.Marshal(v)
}

func (v *WarrantyPromise) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WarrantyPromise"

	return json.Marshal(*v)
}
