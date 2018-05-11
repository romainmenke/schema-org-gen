package schemaorg

import "encoding/json"

// WarrantyPromise see : https://schema.org/WarrantyPromise
type WarrantyPromise struct {

typeContext

StructuredValue

// DurationOfWarranty see : /durationOfWarranty
// The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days.
DurationOfWarranty *QuantitativeValue `json:"durationOfWarranty"`

// WarrantyScope see : /warrantyScope
// The scope of the warranty promise.
WarrantyScope *WarrantyScope `json:"warrantyScope"`

}

func (v *WarrantyPromise) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WarrantyPromise"

	return json.Marshal(v)
}
