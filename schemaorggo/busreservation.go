package schemaorggo

import "encoding/json"

// BusReservation see : https://schema.org/BusReservation
type BusReservation struct {
	Reservation

	typeContext

	// BookingTime see : https://schema.org/bookingTime
	// The date and time the reservation was booked.
	BookingTime DateTime `json:"bookingTime,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	Broker interface{} `json:"broker,omitempty"` // types : Organization Person

	// ModifiedTime see : https://schema.org/modifiedTime
	// The date and time the reservation was modified.
	ModifiedTime DateTime `json:"modifiedTime,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	PriceCurrency string `json:"priceCurrency,omitempty"`

	// ProgramMembershipUsed see : https://schema.org/programMembershipUsed
	// Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
	ProgramMembershipUsed *ProgramMembership `json:"programMembershipUsed,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	Provider interface{} `json:"provider,omitempty"` // types : Organization Person

	// ReservationFor see : https://schema.org/reservationFor
	// The thing -- flight, event, restaurant,etc. being reserved.
	ReservationFor *Thing `json:"reservationFor,omitempty"`

	// ReservationId see : https://schema.org/reservationId
	// A unique identifier for the reservation.
	ReservationId string `json:"reservationId,omitempty"`

	// ReservationStatus see : https://schema.org/reservationStatus
	// The current status of the reservation.
	ReservationStatus *ReservationStatusType `json:"reservationStatus,omitempty"`

	// ReservedTicket see : https://schema.org/reservedTicket
	// A ticket associated with the reservation.
	ReservedTicket *Ticket `json:"reservedTicket,omitempty"`

	// TotalPrice see : https://schema.org/totalPrice
	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	TotalPrice interface{} `json:"totalPrice,omitempty"` // types : Number PriceSpecification Text

	// UnderName see : https://schema.org/underName
	// The person or organization the reservation or ticket is for.
	UnderName interface{} `json:"underName,omitempty"` // types : Organization Person

}

func (v BusReservation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusReservation"

	return json.Marshal(v)
}

func (v *BusReservation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BusReservation"

	return json.Marshal(*v)
}
