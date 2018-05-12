package schemaorg

import "encoding/json"

// UnitPriceSpecification see : https://schema.org/UnitPriceSpecification
type UnitPriceSpecification struct {

PriceSpecification

typeContext

// BillingIncrement see : https://schema.org/billingIncrement
// This property specifies the minimal quantity and rounding increment that will be the basis for the billing. The unit of measurement is specified by the unitCode property.
BillingIncrement float64 `json:"billingIncrement"`

// PriceType see : https://schema.org/priceType
// A short text or acronym indicating multiple price specifications for the same offer, e.g. SRP for the suggested retail price or INVOICE for the invoice price, mostly used in the car industry.
PriceType string `json:"priceType"`

// ReferenceQuantity see : https://schema.org/referenceQuantity
// The reference quantity for which a certain price applies, e.g. 1 EUR per 4 kWh of electricity. This property is a replacement for unitOfMeasurement for the advanced cases where the price does not relate to a standard unit.
ReferenceQuantity *QuantitativeValue `json:"referenceQuantity"`

// UnitCode see : https://schema.org/unitCode
// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
UnitCode interface{} `json:"unitCode"` // types : Text URL

// UnitText see : https://schema.org/unitText
// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
// unitCode (see: https://schema.orgunitCode).
UnitText string `json:"unitText"`

}

func (v *UnitPriceSpecification) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "UnitPriceSpecification"

	return json.Marshal(v)
}
