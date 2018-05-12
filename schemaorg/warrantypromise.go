package schemaorg

import "encoding/json"

// WarrantyPromise see : https://schema.org/WarrantyPromise
type WarrantyPromise struct {

StructuredValue

typeContext

// DurationOfWarranty see : https://schema.org/durationOfWarranty
// The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days.
DurationOfWarranty *QuantitativeValue `json:"durationOfWarranty"`

// WarrantyScope see : https://schema.org/warrantyScope
// The scope of the warranty promise.
WarrantyScope *WarrantyScope `json:"warrantyScope"`

}

func (v *WarrantyPromise) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WarrantyPromise"

	return json.Marshal(v)
}
