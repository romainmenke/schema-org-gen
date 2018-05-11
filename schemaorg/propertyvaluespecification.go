package schemaorg

import "encoding/json"

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
type PropertyValueSpecification struct {

typeContext

Intangible

// DefaultValue see : /defaultValue
// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it's an ID reference to one of the current values.
DefaultValue interface{} `json:"defaultValue"` // types : Text Thing

// MaxValue see : /maxValue
// The upper value of some characteristic or property.
MaxValue float64 `json:"maxValue"`

// MinValue see : /minValue
// The lower value of some characteristic or property.
MinValue float64 `json:"minValue"`

// MultipleValues see : /multipleValues
// Whether multiple values are allowed for the property.  Default is false.
MultipleValues bool `json:"multipleValues"`

// ReadonlyValue see : /readonlyValue
// Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a "hidden" input in an HTML form.
ReadonlyValue bool `json:"readonlyValue"`

// StepValue see : /stepValue
// The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
StepValue float64 `json:"stepValue"`

// ValueMaxLength see : /valueMaxLength
// Specifies the allowed range for number of characters in a literal value.
ValueMaxLength float64 `json:"valueMaxLength"`

// ValueMinLength see : /valueMinLength
// Specifies the minimum allowed range for number of characters in a literal value.
ValueMinLength float64 `json:"valueMinLength"`

// ValueName see : /valueName
// Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML's input@name.
ValueName string `json:"valueName"`

// ValuePattern see : /valuePattern
// Specifies a regular expression for testing literal values according to the HTML spec.
ValuePattern string `json:"valuePattern"`

// ValueRequired see : /valueRequired
// Whether the property must be filled in to complete the action.  Default is false.
ValueRequired bool `json:"valueRequired"`

}

func (v *PropertyValueSpecification) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PropertyValueSpecification"

	return json.Marshal(v)
}
