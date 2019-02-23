package schemaorggo

import "encoding/json"

// DriveWheelConfigurationValue see : https://schema.org/DriveWheelConfigurationValue
type DriveWheelConfigurationValue struct {

	// With properties from QualitativeValue see : https://schema.org/QualitativeValue
	//

	// With properties from Enumeration see : https://schema.org/Enumeration
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalProperty see : https://schema.org/additionalProperty
	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	//
	// types : PropertyValue
	AdditionalProperty []*PropertyValue `json:"additionalProperty,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Equal see : https://schema.org/equal
	// This ordering relation for qualitative values indicates that the subject is equal to the object.
	// types : QualitativeValue
	Equal []*QualitativeValue `json:"equal,omitempty"`

	// Greater see : https://schema.org/greater
	// This ordering relation for qualitative values indicates that the subject is greater than the object.
	// types : QualitativeValue
	Greater []*QualitativeValue `json:"greater,omitempty"`

	// GreaterOrEqual see : https://schema.org/greaterOrEqual
	// This ordering relation for qualitative values indicates that the subject is greater than or equal to the object.
	// types : QualitativeValue
	GreaterOrEqual []*QualitativeValue `json:"greaterOrEqual,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// Lesser see : https://schema.org/lesser
	// This ordering relation for qualitative values indicates that the subject is lesser than the object.
	// types : QualitativeValue
	Lesser []*QualitativeValue `json:"lesser,omitempty"`

	// LesserOrEqual see : https://schema.org/lesserOrEqual
	// This ordering relation for qualitative values indicates that the subject is lesser than or equal to the object.
	// types : QualitativeValue
	LesserOrEqual []*QualitativeValue `json:"lesserOrEqual,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NonEqual see : https://schema.org/nonEqual
	// This ordering relation for qualitative values indicates that the subject is not equal to the object.
	// types : QualitativeValue
	NonEqual []*QualitativeValue `json:"nonEqual,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// ValueReference see : https://schema.org/valueReference
	// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	// types : Enumeration StructuredValue PropertyValue QualitativeValue QuantitativeValue
	ValueReference []interface{} `json:"valueReference,omitempty"`
}

func (v DriveWheelConfigurationValue) MarshalJSON() ([]byte, error) {
	type Alias DriveWheelConfigurationValue

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"DriveWheelConfigurationValue\","), b[1:]...), nil
}
