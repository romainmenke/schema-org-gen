package schemaorg

import "encoding/json"

// Quantity see : https://schema.org/Quantity
type Quantity struct {

Intangible

typeContext

// AdditionalType see : https://schema.org/additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : https://schema.org/alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : https://schema.org/description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : https://schema.org/identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : https://schema.org/image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : https://schema.org/mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : https://schema.org/name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : https://schema.org/potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction *Action `json:"potentialAction"`

// SameAs see : https://schema.org/sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : https://schema.org/url
// URL of the item.
Url string `json:"url"`

}

func (v *Quantity) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Quantity"

	return json.Marshal(v)
}
