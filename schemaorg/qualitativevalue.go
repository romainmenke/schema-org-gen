package schemaorg

import "encoding/json"

// QualitativeValue see : https://schema.org/QualitativeValue
type QualitativeValue struct {

Enumeration

typeContext

// AdditionalProperty see : https://schema.org/additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty *PropertyValue `json:"additionalProperty"`

// Equal see : https://schema.org/equal
// This ordering relation for qualitative values indicates that the subject is equal to the object.
Equal *QualitativeValue `json:"equal"`

// Greater see : https://schema.org/greater
// This ordering relation for qualitative values indicates that the subject is greater than the object.
Greater *QualitativeValue `json:"greater"`

// GreaterOrEqual see : https://schema.org/greaterOrEqual
// This ordering relation for qualitative values indicates that the subject is greater than or equal to the object.
GreaterOrEqual *QualitativeValue `json:"greaterOrEqual"`

// Lesser see : https://schema.org/lesser
// This ordering relation for qualitative values indicates that the subject is lesser than the object.
Lesser *QualitativeValue `json:"lesser"`

// LesserOrEqual see : https://schema.org/lesserOrEqual
// This ordering relation for qualitative values indicates that the subject is lesser than or equal to the object.
LesserOrEqual *QualitativeValue `json:"lesserOrEqual"`

// NonEqual see : https://schema.org/nonEqual
// This ordering relation for qualitative values indicates that the subject is not equal to the object.
NonEqual *QualitativeValue `json:"nonEqual"`

// ValueReference see : https://schema.org/valueReference
// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
ValueReference interface{} `json:"valueReference"` // types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue

}

func (v *QualitativeValue) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "QualitativeValue"

	return json.Marshal(v)
}
