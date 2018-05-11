package schemaorg

import "encoding/json"

// MonetaryAmount see : https://schema.org/MonetaryAmount
type MonetaryAmount struct {

typeContext

StructuredValue

// Currency see : /currency
// The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
Currency string `json:"currency"`

// MaxValue see : /maxValue
// The upper value of some characteristic or property.
MaxValue float64 `json:"maxValue"`

// MinValue see : /minValue
// The lower value of some characteristic or property.
MinValue float64 `json:"minValue"`

// ValidFrom see : /validFrom
// The date when the item becomes valid.
ValidFrom interface{} `json:"validFrom"`

// ValidThrough see : /validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough interface{} `json:"validThrough"`

// Value see : /value
// The value of the quantitative value or property value node.
// 
// 
// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is 'Number'.
// For PropertyValue (see: https://schema.org/PropertyValue), it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.
// 
// 
Value interface{} `json:"value"` // types : Boolean Number StructuredValue Text

}

func (v *MonetaryAmount) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MonetaryAmount"

	return json.Marshal(v)
}
