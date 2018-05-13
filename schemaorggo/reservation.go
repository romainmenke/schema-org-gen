package schemaorggo

import "encoding/json"

// Reservation see : https://schema.org/Reservation
type Reservation struct {
	Intangible

	typeContext

	// BookingTime see : https://schema.org/bookingTime
	// The date and time the reservation was booked.
	// types : DateTime
	BookingTime []DateTime `json:"bookingTime,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	// types : Organization Person
	Broker []interface{} `json:"broker,omitempty"`

	// ModifiedTime see : https://schema.org/modifiedTime
	// The date and time the reservation was modified.
	// types : DateTime
	ModifiedTime []DateTime `json:"modifiedTime,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	// types : Text
	PriceCurrency []string `json:"priceCurrency,omitempty"`

	// ProgramMembershipUsed see : https://schema.org/programMembershipUsed
	// Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
	// types : ProgramMembership
	ProgramMembershipUsed []*ProgramMembership `json:"programMembershipUsed,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`

	// ReservationFor see : https://schema.org/reservationFor
	// The thing -- flight, event, restaurant,etc. being reserved.
	// types : Thing
	ReservationFor []*Thing `json:"reservationFor,omitempty"`

	// ReservationId see : https://schema.org/reservationId
	// A unique identifier for the reservation.
	// types : Text
	ReservationId []string `json:"reservationId,omitempty"`

	// ReservationStatus see : https://schema.org/reservationStatus
	// The current status of the reservation.
	// types : ReservationStatusType
	ReservationStatus []*ReservationStatusType `json:"reservationStatus,omitempty"`

	// ReservedTicket see : https://schema.org/reservedTicket
	// A ticket associated with the reservation.
	// types : Ticket
	ReservedTicket []*Ticket `json:"reservedTicket,omitempty"`

	// TotalPrice see : https://schema.org/totalPrice
	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	// types : Number PriceSpecification Text
	TotalPrice []interface{} `json:"totalPrice,omitempty"`

	// UnderName see : https://schema.org/underName
	// The person or organization the reservation or ticket is for.
	// types : Organization Person
	UnderName []interface{} `json:"underName,omitempty"`
}

func (v Reservation) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BookingTime
		if len(v.BookingTime) == 1 {
			value = v.BookingTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bookingTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Broker
		if len(v.Broker) == 1 {
			value = v.Broker[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broker"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ModifiedTime
		if len(v.ModifiedTime) == 1 {
			value = v.ModifiedTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["modifiedTime"] = json.RawMessage(b)
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
		var value interface{} = v.ProgramMembershipUsed
		if len(v.ProgramMembershipUsed) == 1 {
			value = v.ProgramMembershipUsed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["programMembershipUsed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Provider
		if len(v.Provider) == 1 {
			value = v.Provider[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["provider"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReservationFor
		if len(v.ReservationFor) == 1 {
			value = v.ReservationFor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reservationFor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReservationId
		if len(v.ReservationId) == 1 {
			value = v.ReservationId[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reservationId"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReservationStatus
		if len(v.ReservationStatus) == 1 {
			value = v.ReservationStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reservationStatus"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReservedTicket
		if len(v.ReservedTicket) == 1 {
			value = v.ReservedTicket[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reservedTicket"] = json.RawMessage(b)
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

func (v Reservation) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Reservation"

	return data, nil
}

func (v Reservation) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
