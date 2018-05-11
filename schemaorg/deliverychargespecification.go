package schemaorg

// DeliveryChargeSpecification see : https://schema.org/DeliveryChargeSpecification
type DeliveryChargeSpecification struct {

PriceSpecification// AppliesToDeliveryMethod see : /appliesToDeliveryMethod
// The delivery method(s) to which the delivery charge or payment charge specification applies.
AppliesToDeliveryMethod string `json:"appliesToDeliveryMethod"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// EligibleRegion see : /eligibleRegion
// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
// 
// See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
EligibleRegion interface{} `json:"eligibleRegion"` // types : GeoShape Place Text

// IneligibleRegion see : /ineligibleRegion
// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
// 
// See also eligibleRegion (see: https://schema.org/eligibleRegion).
IneligibleRegion interface{} `json:"ineligibleRegion"` // types : GeoShape Place Text

// EligibleQuantity see : /eligibleQuantity
// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
EligibleQuantity string `json:"eligibleQuantity"`

// EligibleTransactionVolume see : /eligibleTransactionVolume
// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
EligibleTransactionVolume string `json:"eligibleTransactionVolume"`

// MaxPrice see : /maxPrice
// The highest price if the price is a range.
MaxPrice string `json:"maxPrice"`

// MinPrice see : /minPrice
// The lowest price if the price is a range.
MinPrice string `json:"minPrice"`

// Price see : /price
// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
// 
// Usage guidelines:
// 
// 
// Use the priceCurrency (see: https://schema.org/priceCurrency) property (with ISO 4217 codes (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217#Active_codes) e.g. "USD") instead of
//   including ambiguous symbols (see: https://schema.orghttp://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as '$' in the value.
// Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
// Note that both RDFa (see: https://schema.orghttp://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a "content=" attribute for publishing simple machine-readable values alongside more human-friendly formatting.
// Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similiar Unicode symbols.
// 
// 
Price interface{} `json:"price"` // types : Number Text

// PriceCurrency see : /priceCurrency
// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
PriceCurrency string `json:"priceCurrency"`

// ValidFrom see : /validFrom
// The date when the item becomes valid.
ValidFrom string `json:"validFrom"`

// ValidThrough see : /validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough string `json:"validThrough"`

// ValueAddedTaxIncluded see : /valueAddedTaxIncluded
// Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
ValueAddedTaxIncluded string `json:"valueAddedTaxIncluded"`

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

