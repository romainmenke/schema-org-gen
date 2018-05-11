package schemaorg

// SteeringPositionValue see : https://schema.org/SteeringPositionValue
type SteeringPositionValue struct {

QualitativeValue// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty string `json:"additionalProperty"`

// Equal see : /equal
// This ordering relation for qualitative values indicates that the subject is equal to the object.
Equal string `json:"equal"`

// Greater see : /greater
// This ordering relation for qualitative values indicates that the subject is greater than the object.
Greater string `json:"greater"`

// GreaterOrEqual see : /greaterOrEqual
// This ordering relation for qualitative values indicates that the subject is greater than or equal to the object.
GreaterOrEqual string `json:"greaterOrEqual"`

// Lesser see : /lesser
// This ordering relation for qualitative values indicates that the subject is lesser than the object.
Lesser string `json:"lesser"`

// LesserOrEqual see : /lesserOrEqual
// This ordering relation for qualitative values indicates that the subject is lesser than or equal to the object.
LesserOrEqual string `json:"lesserOrEqual"`

// NonEqual see : /nonEqual
// This ordering relation for qualitative values indicates that the subject is not equal to the object.
NonEqual string `json:"nonEqual"`

// ValueReference see : /valueReference
// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
ValueReference interface{} `json:"valueReference"` // types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// SteeringPosition see : /steeringPosition
// The position of the steering wheel or similar device (mostly for cars). 
SteeringPosition string `json:"steeringPosition"`

}

