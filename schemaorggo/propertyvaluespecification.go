package schemaorggo

import "encoding/json"

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
type PropertyValueSpecification struct {
	Intangible

	typeContext

	// DefaultValue see : https://schema.org/defaultValue
	// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it&#39;s an ID reference to one of the current values.
	// types : Text Thing
	DefaultValue []interface{} `json:"defaultValue,omitempty"`

	// MaxValue see : https://schema.org/maxValue
	// The upper value of some characteristic or property.
	// types : Number
	MaxValue []float64 `json:"maxValue,omitempty"`

	// MinValue see : https://schema.org/minValue
	// The lower value of some characteristic or property.
	// types : Number
	MinValue []float64 `json:"minValue,omitempty"`

	// MultipleValues see : https://schema.org/multipleValues
	// Whether multiple values are allowed for the property.  Default is false.
	// types : Boolean
	MultipleValues []bool `json:"multipleValues,omitempty"`

	// ReadonlyValue see : https://schema.org/readonlyValue
	// Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a &quot;hidden&quot; input in an HTML form.
	// types : Boolean
	ReadonlyValue []bool `json:"readonlyValue,omitempty"`

	// StepValue see : https://schema.org/stepValue
	// The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	// types : Number
	StepValue []float64 `json:"stepValue,omitempty"`

	// ValueMaxLength see : https://schema.org/valueMaxLength
	// Specifies the allowed range for number of characters in a literal value.
	// types : Number
	ValueMaxLength []float64 `json:"valueMaxLength,omitempty"`

	// ValueMinLength see : https://schema.org/valueMinLength
	// Specifies the minimum allowed range for number of characters in a literal value.
	// types : Number
	ValueMinLength []float64 `json:"valueMinLength,omitempty"`

	// ValueName see : https://schema.org/valueName
	// Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML&#39;s input@name.
	// types : Text
	ValueName []string `json:"valueName,omitempty"`

	// ValuePattern see : https://schema.org/valuePattern
	// Specifies a regular expression for testing literal values according to the HTML spec.
	// types : Text
	ValuePattern []string `json:"valuePattern,omitempty"`

	// ValueRequired see : https://schema.org/valueRequired
	// Whether the property must be filled in to complete the action.  Default is false.
	// types : Boolean
	ValueRequired []bool `json:"valueRequired,omitempty"`
}

func (v PropertyValueSpecification) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.DefaultValue
		if len(v.DefaultValue) == 1 {
			value = v.DefaultValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["defaultValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MaxValue
		if len(v.MaxValue) == 1 {
			value = v.MaxValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["maxValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MinValue
		if len(v.MinValue) == 1 {
			value = v.MinValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["minValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MultipleValues
		if len(v.MultipleValues) == 1 {
			value = v.MultipleValues[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["multipleValues"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReadonlyValue
		if len(v.ReadonlyValue) == 1 {
			value = v.ReadonlyValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["readonlyValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StepValue
		if len(v.StepValue) == 1 {
			value = v.StepValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["stepValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValueMaxLength
		if len(v.ValueMaxLength) == 1 {
			value = v.ValueMaxLength[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["valueMaxLength"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValueMinLength
		if len(v.ValueMinLength) == 1 {
			value = v.ValueMinLength[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["valueMinLength"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValueName
		if len(v.ValueName) == 1 {
			value = v.ValueName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["valueName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValuePattern
		if len(v.ValuePattern) == 1 {
			value = v.ValuePattern[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["valuePattern"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValueRequired
		if len(v.ValueRequired) == 1 {
			value = v.ValueRequired[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["valueRequired"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PropertyValueSpecification) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PropertyValueSpecification"

	return data, nil
}

func (v PropertyValueSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
