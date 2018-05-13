package schemaorggo

import "encoding/json"

// Ticket see : https://schema.org/Ticket
type Ticket struct {
	Intangible

	typeContext

	// DateIssued see : https://schema.org/dateIssued
	// The date the ticket was issued.
	// types : DateTime
	DateIssued []DateTime `json:"dateIssued,omitempty"`

	// IssuedBy see : https://schema.org/issuedBy
	// The organization issuing the ticket or permit.
	// types : Organization
	IssuedBy []*Organization `json:"issuedBy,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	// types : Text
	PriceCurrency []string `json:"priceCurrency,omitempty"`

	// TicketNumber see : https://schema.org/ticketNumber
	// The unique identifier for the ticket.
	// types : Text
	TicketNumber []string `json:"ticketNumber,omitempty"`

	// TicketToken see : https://schema.org/ticketToken
	// Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
	// types : Text URL
	TicketToken []string `json:"ticketToken,omitempty"`

	// TicketedSeat see : https://schema.org/ticketedSeat
	// The seat associated with the ticket.
	// types : Seat
	TicketedSeat []*Seat `json:"ticketedSeat,omitempty"`

	// TotalPrice see : https://schema.org/totalPrice
	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	// types : Number PriceSpecification Text
	TotalPrice []interface{} `json:"totalPrice,omitempty"`

	// UnderName see : https://schema.org/underName
	// The person or organization the reservation or ticket is for.
	// types : Organization Person
	UnderName []interface{} `json:"underName,omitempty"`
}

func (v Ticket) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.DateIssued
		if len(v.DateIssued) == 1 {
			value = v.DateIssued[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateIssued"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IssuedBy
		if len(v.IssuedBy) == 1 {
			value = v.IssuedBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["issuedBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PriceCurrency
		if len(v.PriceCurrency) == 1 {
			value = v.PriceCurrency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceCurrency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TicketNumber
		if len(v.TicketNumber) == 1 {
			value = v.TicketNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ticketNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TicketToken
		if len(v.TicketToken) == 1 {
			value = v.TicketToken[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ticketToken"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TicketedSeat
		if len(v.TicketedSeat) == 1 {
			value = v.TicketedSeat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ticketedSeat"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TotalPrice
		if len(v.TotalPrice) == 1 {
			value = v.TotalPrice[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["totalPrice"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UnderName
		if len(v.UnderName) == 1 {
			value = v.UnderName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["underName"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Ticket) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Ticket"

	return data, nil
}

func (v Ticket) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
