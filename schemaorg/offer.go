package schemaorg

// Offer see : https://schema.org/Offer
type Offer struct {

Intangible// AcceptedPaymentMethod see : /acceptedPaymentMethod
// The payment method(s) accepted by seller for this offer.
AcceptedPaymentMethod interface{} `json:"acceptedPaymentMethod"` // types : LoanOrCredit PaymentMethod

// AddOn see : /addOn
// An additional offer that can only be obtained in combination with the first base offer (e.g. supplements and extensions that are available for a surcharge).
AddOn string `json:"addOn"`

// AdvanceBookingRequirement see : /advanceBookingRequirement
// The amount of time that is required between accepting the offer and the actual usage of the resource or service.
AdvanceBookingRequirement string `json:"advanceBookingRequirement"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Availability see : /availability
// The availability of this item—for example In stock, Out of stock, Pre-order, etc.
Availability string `json:"availability"`

// AvailabilityEnds see : /availabilityEnds
// The end of the availability of the product or service included in the offer.
AvailabilityEnds string `json:"availabilityEnds"`

// AvailabilityStarts see : /availabilityStarts
// The beginning of the availability of the product or service included in the offer.
AvailabilityStarts string `json:"availabilityStarts"`

// AvailableAtOrFrom see : /availableAtOrFrom
// The place(s) from which the offer can be obtained (e.g. store locations).
AvailableAtOrFrom string `json:"availableAtOrFrom"`

// AvailableDeliveryMethod see : /availableDeliveryMethod
// The delivery method(s) available for this offer.
AvailableDeliveryMethod string `json:"availableDeliveryMethod"`

// BusinessFunction see : /businessFunction
// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
BusinessFunction string `json:"businessFunction"`

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
Category interface{} `json:"category"` // types : PhysicalActivityCategory Text Thing

// DeliveryLeadTime see : /deliveryLeadTime
// The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup.
DeliveryLeadTime string `json:"deliveryLeadTime"`

// EligibleCustomerType see : /eligibleCustomerType
// The type(s) of customers for which the given offer is valid.
EligibleCustomerType string `json:"eligibleCustomerType"`

// EligibleDuration see : /eligibleDuration
// The duration for which the given offer is valid.
EligibleDuration string `json:"eligibleDuration"`

// EligibleQuantity see : /eligibleQuantity
// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
EligibleQuantity string `json:"eligibleQuantity"`

// EligibleRegion see : /eligibleRegion
// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
// 
// See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
EligibleRegion interface{} `json:"eligibleRegion"` // types : GeoShape Place Text

// EligibleTransactionVolume see : /eligibleTransactionVolume
// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
EligibleTransactionVolume string `json:"eligibleTransactionVolume"`

// Gtin12 see : /gtin12
// The GTIN-12 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx) code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin12 string `json:"gtin12"`

// Gtin13 see : /gtin13
// The GTIN-13 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx) code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin13 string `json:"gtin13"`

// Gtin14 see : /gtin14
// The GTIN-14 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx) code of the product, or the product to which the offer refers. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin14 string `json:"gtin14"`

// Gtin8 see : /gtin8
// The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin8 string `json:"gtin8"`

// IncludesObject see : /includesObject
// This links to a node or nodes indicating the exact quantity of the products included in the offer.
IncludesObject string `json:"includesObject"`

// IneligibleRegion see : /ineligibleRegion
// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
// 
// See also eligibleRegion (see: https://schema.org/eligibleRegion).
IneligibleRegion interface{} `json:"ineligibleRegion"` // types : GeoShape Place Text

// InventoryLevel see : /inventoryLevel
// The current approximate inventory level for the item or items.
InventoryLevel string `json:"inventoryLevel"`

// ItemCondition see : /itemCondition
// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
ItemCondition string `json:"itemCondition"`

// ItemOffered see : /itemOffered
// The item being offered.
ItemOffered interface{} `json:"itemOffered"` // types : Product Service

// Mpn see : /mpn
// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
Mpn string `json:"mpn"`

// OfferedBy see : /offeredBy
// A pointer to the organization or person making the offer. Inverse property: makesOffer (see: https://schema.org/makesOffer).
OfferedBy interface{} `json:"offeredBy"` // types : Organization Person

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

// PriceSpecification see : /priceSpecification
// One or more detailed price specifications, indicating the unit price and delivery or payment charges.
PriceSpecification string `json:"priceSpecification"`

// PriceValidUntil see : /priceValidUntil
// The date after which the price is no longer available.
PriceValidUntil string `json:"priceValidUntil"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// Seller see : /seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
Seller interface{} `json:"seller"` // types : Organization Person

// SerialNumber see : /serialNumber
// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
SerialNumber string `json:"serialNumber"`

// Sku see : /sku
// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
Sku string `json:"sku"`

// ValidFrom see : /validFrom
// The date when the item becomes valid.
ValidFrom string `json:"validFrom"`

// ValidThrough see : /validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough string `json:"validThrough"`

// Warranty see : /warranty
// The warranty promise(s) included in the offer. Supersedes warrantyPromise (see: https://schema.org/warrantyPromise).
Warranty string `json:"warranty"`

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

// AcceptedOffer see : /acceptedOffer
// The offer(s) -- e.g., product, quantity and price combinations -- included in the order. 
AcceptedOffer string `json:"acceptedOffer"`

// AddOn see : /addOn
// An additional offer that can only be obtained in combination with the first base offer (e.g. supplements and extensions that are available for a surcharge). 
AddOn string `json:"addOn"`

// ExpectsAcceptanceOf see : /expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it. 
ExpectsAcceptanceOf string `json:"expectsAcceptanceOf"`

// MakesOffer see : /makesOffer
// A pointer to products or services offered by the organization or person.  inverse property: offeredBy (see: https://schema.org/offeredBy).
MakesOffer interface{} `json:"makesOffer"` // types : Organization Person

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event. 
Offers interface{} `json:"offers"` // types : AggregateOffer CreativeWork Event MenuItem Product Service

}

