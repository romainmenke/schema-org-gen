package schemaorg

import "encoding/json"

// Reservation see : https://schema.org/Reservation
type Reservation struct {

typeContext

Intangible

// BookingTime see : /bookingTime
// The date and time the reservation was booked.
BookingTime interface{} `json:"bookingTime"`

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Organization Person

// ModifiedTime see : /modifiedTime
// The date and time the reservation was modified.
ModifiedTime interface{} `json:"modifiedTime"`

// PriceCurrency see : /priceCurrency
// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
PriceCurrency string `json:"priceCurrency"`

// ProgramMembershipUsed see : /programMembershipUsed
// Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
ProgramMembershipUsed *ProgramMembership `json:"programMembershipUsed"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// ReservationFor see : /reservationFor
// The thing -- flight, event, restaurant,etc. being reserved.
ReservationFor *Thing `json:"reservationFor"`

// ReservationId see : /reservationId
// A unique identifier for the reservation.
ReservationId string `json:"reservationId"`

// ReservationStatus see : /reservationStatus
// The current status of the reservation.
ReservationStatus *ReservationStatusType `json:"reservationStatus"`

// ReservedTicket see : /reservedTicket
// A ticket associated with the reservation.
ReservedTicket *Ticket `json:"reservedTicket"`

// TotalPrice see : /totalPrice
// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
TotalPrice interface{} `json:"totalPrice"` // types : Number PriceSpecification Text

// UnderName see : /underName
// The person or organization the reservation or ticket is for.
UnderName interface{} `json:"underName"` // types : Organization Person

}

func (v *Reservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Reservation"

	return json.Marshal(v)
}
