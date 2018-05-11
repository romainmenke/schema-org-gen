package schemaorg

// Ticket see : https://schema.org/Ticket
type Ticket struct {

Intangible// DateIssued see : /dateIssued
// The date the ticket was issued.
DateIssued string `json:"dateIssued"`

// IssuedBy see : /issuedBy
// The organization issuing the ticket or permit.
IssuedBy string `json:"issuedBy"`

// PriceCurrency see : /priceCurrency
// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
PriceCurrency string `json:"priceCurrency"`

// TicketNumber see : /ticketNumber
// The unique identifier for the ticket.
TicketNumber string `json:"ticketNumber"`

// TicketToken see : /ticketToken
// Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
TicketToken interface{} `json:"ticketToken"` // types : Text URL

// TicketedSeat see : /ticketedSeat
// The seat associated with the ticket.
TicketedSeat string `json:"ticketedSeat"`

// TotalPrice see : /totalPrice
// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
TotalPrice interface{} `json:"totalPrice"` // types : Number PriceSpecification Text

// UnderName see : /underName
// The person or organization the reservation or ticket is for.
UnderName interface{} `json:"underName"` // types : Organization Person

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

// ReservedTicket see : /reservedTicket
// A ticket associated with the reservation. 
ReservedTicket string `json:"reservedTicket"`

}

