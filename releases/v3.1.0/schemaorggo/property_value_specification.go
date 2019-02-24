package schemaorggo

import "encoding/json"

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
type PropertyValueSpecification struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// DefaultValue see : https://schema.org/defaultValue
	// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it&#39;s an ID reference to one of the current values.
	// types : Thing Text
	DefaultValue []interface{} `json:"defaultValue,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

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

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// ReadonlyValue see : https://schema.org/readonlyValue
	// Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a &quot;hidden&quot; input in an HTML form.
	// types : Boolean
	ReadonlyValue []bool `json:"readonlyValue,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// StepValue see : https://schema.org/stepValue
	// The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	// types : Number
	StepValue []float64 `json:"stepValue,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

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

func (v PropertyValueSpecification) MarshalJSON() ([]byte, error) {
	type Alias PropertyValueSpecification

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"PropertyValueSpecification\","), b[1:]...), nil
}
