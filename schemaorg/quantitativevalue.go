package schemaorg

import "encoding/json"

// QuantitativeValue see : https://schema.org/QuantitativeValue
type QuantitativeValue struct {

typeContext

StructuredValue

// AdditionalProperty see : https://schema.org/additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty *PropertyValue `json:"additionalProperty"`

// MaxValue see : https://schema.org/maxValue
// The upper value of some characteristic or property.
MaxValue float64 `json:"maxValue"`

// MinValue see : https://schema.org/minValue
// The lower value of some characteristic or property.
MinValue float64 `json:"minValue"`

// UnitCode see : https://schema.org/unitCode
// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
UnitCode interface{} `json:"unitCode"` // types : Text URL

// UnitText see : https://schema.org/unitText
// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
// unitCode (see: https://schema.orgunitCode).
UnitText string `json:"unitText"`

// Value see : https://schema.org/value
// The value of the quantitative value or property value node.
// 
// 
// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is 'Number'.
// For PropertyValue (see: https://schema.org/PropertyValue), it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.
// 
// 
Value interface{} `json:"value"` // types : Boolean Number StructuredValue Text

// ValueReference see : https://schema.org/valueReference
// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
ValueReference interface{} `json:"valueReference"` // types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue

}

func (v *QuantitativeValue) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "QuantitativeValue"

	return json.Marshal(v)
}
