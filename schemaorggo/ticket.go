package schemaorggo

import "encoding/json"

// Ticket see : https://schema.org/Ticket
type Ticket struct {
	Intangible

	typeContext

	// DateIssued see : https://schema.org/dateIssued
	// The date the ticket was issued.
	DateIssued DateTime `json:"dateIssued,omitempty"` // types : DateTime

	// IssuedBy see : https://schema.org/issuedBy
	// The organization issuing the ticket or permit.
	IssuedBy *Organization `json:"issuedBy,omitempty"` // types : Organization

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	PriceCurrency string `json:"priceCurrency,omitempty"` // types : Text

	// TicketNumber see : https://schema.org/ticketNumber
	// The unique identifier for the ticket.
	TicketNumber string `json:"ticketNumber,omitempty"` // types : Text

	// TicketToken see : https://schema.org/ticketToken
	// Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
	TicketToken string `json:"ticketToken,omitempty"` // types : Text URL

	// TicketedSeat see : https://schema.org/ticketedSeat
	// The seat associated with the ticket.
	TicketedSeat *Seat `json:"ticketedSeat,omitempty"` // types : Seat

	// TotalPrice see : https://schema.org/totalPrice
	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	TotalPrice interface{} `json:"totalPrice,omitempty"` // types : Number PriceSpecification Text

	// UnderName see : https://schema.org/underName
	// The person or organization the reservation or ticket is for.
	UnderName interface{} `json:"underName,omitempty"` // types : Organization Person

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
