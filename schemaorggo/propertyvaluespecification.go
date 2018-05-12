package schemaorggo

import "encoding/json"

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
type PropertyValueSpecification struct {
	Intangible

	typeContext

	// DefaultValue see : https://schema.org/defaultValue
	// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it&#39;s an ID reference to one of the current values.
	// types : Text Thing
	DefaultValue interface{} `json:"defaultValue,omitempty"`

	// MaxValue see : https://schema.org/maxValue
	// The upper value of some characteristic or property.
	// types : Number
	MaxValue float64 `json:"maxValue,omitempty"`

	// MinValue see : https://schema.org/minValue
	// The lower value of some characteristic or property.
	// types : Number
	MinValue float64 `json:"minValue,omitempty"`

	// MultipleValues see : https://schema.org/multipleValues
	// Whether multiple values are allowed for the property.  Default is false.
	// types : Boolean
	MultipleValues bool `json:"multipleValues,omitempty"`

	// ReadonlyValue see : https://schema.org/readonlyValue
	// Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a &quot;hidden&quot; input in an HTML form.
	// types : Boolean
	ReadonlyValue bool `json:"readonlyValue,omitempty"`

	// StepValue see : https://schema.org/stepValue
	// The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	// types : Number
	StepValue float64 `json:"stepValue,omitempty"`

	// ValueMaxLength see : https://schema.org/valueMaxLength
	// Specifies the allowed range for number of characters in a literal value.
	// types : Number
	ValueMaxLength float64 `json:"valueMaxLength,omitempty"`

	// ValueMinLength see : https://schema.org/valueMinLength
	// Specifies the minimum allowed range for number of characters in a literal value.
	// types : Number
	ValueMinLength float64 `json:"valueMinLength,omitempty"`

	// ValueName see : https://schema.org/valueName
	// Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML&#39;s input@name.
	// types : Text
	ValueName string `json:"valueName,omitempty"`

	// ValuePattern see : https://schema.org/valuePattern
	// Specifies a regular expression for testing literal values according to the HTML spec.
	// types : Text
	ValuePattern string `json:"valuePattern,omitempty"`

	// ValueRequired see : https://schema.org/valueRequired
	// Whether the property must be filled in to complete the action.  Default is false.
	// types : Boolean
	ValueRequired bool `json:"valueRequired,omitempty"`
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
