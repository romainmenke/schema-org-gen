package schemaorggo

import "encoding/json"

// Ticket see : https://schema.org/Ticket
type Ticket struct {
	Intangible

	typeContext

	// DateIssued see : https://schema.org/dateIssued
	// The date the ticket was issued.
	// types : DateTime
	DateIssued DateTime `json:"dateIssued,omitempty"`

	// IssuedBy see : https://schema.org/issuedBy
	// The organization issuing the ticket or permit.
	// types : Organization
	IssuedBy *Organization `json:"issuedBy,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	// types : Text
	PriceCurrency string `json:"priceCurrency,omitempty"`

	// TicketNumber see : https://schema.org/ticketNumber
	// The unique identifier for the ticket.
	// types : Text
	TicketNumber string `json:"ticketNumber,omitempty"`

	// TicketToken see : https://schema.org/ticketToken
	// Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
	// types : Text URL
	TicketToken string `json:"ticketToken,omitempty"`

	// TicketedSeat see : https://schema.org/ticketedSeat
	// The seat associated with the ticket.
	// types : Seat
	TicketedSeat *Seat `json:"ticketedSeat,omitempty"`

	// TotalPrice see : https://schema.org/totalPrice
	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	// types : Number PriceSpecification Text
	TotalPrice interface{} `json:"totalPrice,omitempty"`

	// UnderName see : https://schema.org/underName
	// The person or organization the reservation or ticket is for.
	// types : Organization Person
	UnderName interface{} `json:"underName,omitempty"`
}

func (v Ticket) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Ticket"

	return json.Marshal(v)
}

func (v *Ticket) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Ticket"

	return json.Marshal(*v)
}
