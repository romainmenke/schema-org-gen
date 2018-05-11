package schemaorg

import "encoding/json"

// Ticket see : https://schema.org/Ticket
type Ticket struct {

typeContext

Intangible

// DateIssued see : /dateIssued
// The date the ticket was issued.
DateIssued interface{} `json:"dateIssued"`

// IssuedBy see : /issuedBy
// The organization issuing the ticket or permit.
IssuedBy *Organization `json:"issuedBy"`

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
TicketedSeat *Seat `json:"ticketedSeat"`

// TotalPrice see : /totalPrice
// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
TotalPrice interface{} `json:"totalPrice"` // types : Number PriceSpecification Text

// UnderName see : /underName
// The person or organization the reservation or ticket is for.
UnderName interface{} `json:"underName"` // types : Organization Person

}

func (v *Ticket) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Ticket"

	return json.Marshal(v)
}
