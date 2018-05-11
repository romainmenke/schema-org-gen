package schemaorg

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {

StructuredValue// Calories see : /calories
// The number of calories.
Calories string `json:"calories"`

// CarbohydrateContent see : /carbohydrateContent
// The number of grams of carbohydrates.
CarbohydrateContent string `json:"carbohydrateContent"`

// CholesterolContent see : /cholesterolContent
// The number of milligrams of cholesterol.
CholesterolContent string `json:"cholesterolContent"`

// FatContent see : /fatContent
// The number of grams of fat.
FatContent string `json:"fatContent"`

// FiberContent see : /fiberContent
// The number of grams of fiber.
FiberContent string `json:"fiberContent"`

// ProteinContent see : /proteinContent
// The number of grams of protein.
ProteinContent string `json:"proteinContent"`

// SaturatedFatContent see : /saturatedFatContent
// The number of grams of saturated fat.
SaturatedFatContent string `json:"saturatedFatContent"`

// ServingSize see : /servingSize
// The serving size, in terms of the number of volume or mass.
ServingSize string `json:"servingSize"`

// SodiumContent see : /sodiumContent
// The number of milligrams of sodium.
SodiumContent string `json:"sodiumContent"`

// SugarContent see : /sugarContent
// The number of grams of sugar.
SugarContent string `json:"sugarContent"`

// TransFatContent see : /transFatContent
// The number of grams of trans fat.
TransFatContent string `json:"transFatContent"`

// UnsaturatedFatContent see : /unsaturatedFatContent
// The number of grams of unsaturated fat.
UnsaturatedFatContent string `json:"unsaturatedFatContent"`

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

// Nutrition see : /nutrition
// Nutrition information about the recipe or menu item. 
Nutrition interface{} `json:"nutrition"` // types : MenuItem Recipe

}

