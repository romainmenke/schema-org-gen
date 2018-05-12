package schemaorggo

import "encoding/json"

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
type PropertyValueSpecification struct {
	Intangible

	typeContext

	// DefaultValue see : https://schema.org/defaultValue
	// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it's an ID reference to one of the current values.
	DefaultValue interface{} `json:"defaultValue,omitempty"` // types : Text Thing

	// MaxValue see : https://schema.org/maxValue
	// The upper value of some characteristic or property.
	MaxValue float64 `json:"maxValue,omitempty"` // types : Number

	// MinValue see : https://schema.org/minValue
	// The lower value of some characteristic or property.
	MinValue float64 `json:"minValue,omitempty"` // types : Number

	// MultipleValues see : https://schema.org/multipleValues
	// Whether multiple values are allowed for the property.  Default is false.
	MultipleValues bool `json:"multipleValues,omitempty"` // types : Boolean

	// ReadonlyValue see : https://schema.org/readonlyValue
	// Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a "hidden" input in an HTML form.
	ReadonlyValue bool `json:"readonlyValue,omitempty"` // types : Boolean

	// StepValue see : https://schema.org/stepValue
	// The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	StepValue float64 `json:"stepValue,omitempty"` // types : Number

	// ValueMaxLength see : https://schema.org/valueMaxLength
	// Specifies the allowed range for number of characters in a literal value.
	ValueMaxLength float64 `json:"valueMaxLength,omitempty"` // types : Number

	// ValueMinLength see : https://schema.org/valueMinLength
	// Specifies the minimum allowed range for number of characters in a literal value.
	ValueMinLength float64 `json:"valueMinLength,omitempty"` // types : Number

	// ValueName see : https://schema.org/valueName
	// Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML's input@name.
	ValueName string `json:"valueName,omitempty"` // types : Text

	// ValuePattern see : https://schema.org/valuePattern
	// Specifies a regular expression for testing literal values according to the HTML spec.
	ValuePattern string `json:"valuePattern,omitempty"` // types : Text

	// ValueRequired see : https://schema.org/valueRequired
	// Whether the property must be filled in to complete the action.  Default is false.
	ValueRequired bool `json:"valueRequired,omitempty"` // types : Boolean

}

func (v PropertyValueSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PropertyValueSpecification"

	return json.Marshal(v)
}

func (v *PropertyValueSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PropertyValueSpecification"

	return json.Marshal(*v)
}
