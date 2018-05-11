package schemaorg

// ParcelDelivery see : https://schema.org/ParcelDelivery
type ParcelDelivery struct {

Intangible// DeliveryAddress see : /deliveryAddress
// Destination address.
DeliveryAddress string `json:"deliveryAddress"`

// DeliveryStatus see : /deliveryStatus
// New entry added as the package passes through each leg of its journey (from shipment to final delivery).
DeliveryStatus string `json:"deliveryStatus"`

// ExpectedArrivalFrom see : /expectedArrivalFrom
// The earliest date the package may arrive.
ExpectedArrivalFrom string `json:"expectedArrivalFrom"`

// ExpectedArrivalUntil see : /expectedArrivalUntil
// The latest date the package may arrive.
ExpectedArrivalUntil string `json:"expectedArrivalUntil"`

// HasDeliveryMethod see : /hasDeliveryMethod
// Method used for delivery or shipping.
HasDeliveryMethod string `json:"hasDeliveryMethod"`

// ItemShipped see : /itemShipped
// Item(s) being shipped.
ItemShipped string `json:"itemShipped"`

// OriginAddress see : /originAddress
// Shipper's address.
OriginAddress string `json:"originAddress"`

// PartOfOrder see : /partOfOrder
// The overall order the items in this delivery were included in.
PartOfOrder string `json:"partOfOrder"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// TrackingNumber see : /trackingNumber
// Shipper tracking number.
TrackingNumber string `json:"trackingNumber"`

// TrackingUrl see : /trackingUrl
// Tracking url for the parcel delivery.
TrackingUrl string `json:"trackingUrl"`

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

// OrderDelivery see : /orderDelivery
// The delivery of the parcel related to this order or order item. 
OrderDelivery interface{} `json:"orderDelivery"` // types : Order OrderItem

}

