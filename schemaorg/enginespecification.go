package schemaorg

// EngineSpecification see : https://schema.org/EngineSpecification
type EngineSpecification struct {

StructuredValue// EngineDisplacement see : http://auto.schema.org/engineDisplacement
// The volume swept by all of the pistons inside the cylinders of an internal combustion engine in a single movement. 
// 
// Typical unit code(s): CMQ for cubic centimeter, LTR for liters, INQ for cubic inches
// * Note 1: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
// * Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
EngineDisplacement string `json:"engineDisplacement"`

// EnginePower see : http://auto.schema.org/enginePower
// The power of the vehicle's engine.
//     Typical unit code(s): KWT for kilowatt, BHP for brake horsepower, N12 for metric horsepower (PS, with 1 PS = 735,49875 W)
// 
// 
// Note 1: There are many different ways of measuring an engine's power. For an overview, see  http://en.wikipedia.org/wiki/Horsepower#Engine (see: https://schema.orghttp://en.wikipedia.org/wiki/Horsepower#Engine_power_test_codes)powertest_codes.
// Note 2: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
// Note 3: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
// 
// 
EnginePower string `json:"enginePower"`

// EngineType see : http://auto.schema.org/engineType
// The type of engine or engines powering the vehicle.
EngineType interface{} `json:"engineType"` // types : QualitativeValue Text URL

// FuelType see : /fuelType
// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
FuelType interface{} `json:"fuelType"` // types : QualitativeValue Text URL

// Torque see : http://auto.schema.org/torque
// The torque (turning force) of the vehicle's engine.
// 
// Typical unit code(s): NU for newton metre (N m), F17 for pound-force per foot, or F48 for pound-force per inch
// 
// 
// Note 1: You can link to information about how the given value has been determined (e.g. reference RPM) using the valueReference (see: https://schema.org/valueReference) property.
// Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
// 
// 
Torque string `json:"torque"`

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

// VehicleEngine see : /vehicleEngine
// Information about the engine or engines of the vehicle. 
VehicleEngine string `json:"vehicleEngine"`

}

