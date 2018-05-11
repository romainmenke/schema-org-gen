package schemaorg

// HowToDirection see : https://schema.org/HowToDirection
type HowToDirection struct {

ListItem// AfterMedia see : /afterMedia
// A media object representing the circumstances after performing this direction.
AfterMedia string `json:"afterMedia"`

// BeforeMedia see : /beforeMedia
// A media object representing the circumstances before performing this direction.
BeforeMedia string `json:"beforeMedia"`

// DuringMedia see : /duringMedia
// A media object representing the circumstances while performing this direction.
DuringMedia string `json:"duringMedia"`

// PerformTime see : /performTime
// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
PerformTime string `json:"performTime"`

// PrepTime see : /prepTime
// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
PrepTime string `json:"prepTime"`

// Supply see : /supply
// A sub-property of instrument. A supply consumed when performing instructions or a direction.
Supply interface{} `json:"supply"` // types : HowToSupply Text

// Tool see : /tool
// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
Tool interface{} `json:"tool"` // types : HowToTool Text

// TotalTime see : /totalTime
// The total time required to perform instructions or a direction (including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
TotalTime string `json:"totalTime"`

// Item see : /item
// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
Item string `json:"item"`

// NextItem see : /nextItem
// A link to the ListItem that follows the current one.
NextItem string `json:"nextItem"`

// Position see : /position
// The position of an item in a series or sequence of items.
Position interface{} `json:"position"` // types : Integer Text

// PreviousItem see : /previousItem
// A link to the ListItem that preceeds the current one.
PreviousItem string `json:"previousItem"`

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

}

