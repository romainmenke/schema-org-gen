package schemaorggo

import "encoding/json"

// FoodEstablishmentReservation see : https://schema.org/FoodEstablishmentReservation
type FoodEstablishmentReservation struct {

	// With properties from Reservation see : https://schema.org/Reservation
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// BookingAgent see : https://schema.org/bookingAgent
	// &#39;bookingAgent&#39; is an out-dated term indicating a &#39;broker&#39; that serves as a booking agent.
	// types : Person Organization
	BookingAgent []interface{} `json:"bookingAgent,omitempty"`

	// BookingTime see : https://schema.org/bookingTime
	// The date and time the reservation was booked.
	// types : DateTime
	BookingTime []DateTime `json:"bookingTime,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	// types : Person Organization
	Broker []interface{} `json:"broker,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// EndTime see : https://schema.org/endTime
	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to *December*.\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	EndTime []DateTime `json:"endTime,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// ModifiedTime see : https://schema.org/modifiedTime
	// The date and time the reservation was modified.
	// types : DateTime
	ModifiedTime []DateTime `json:"modifiedTime,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PartySize see : https://schema.org/partySize
	// Number of people the reservation should accommodate.
	// types : Integer QuantitativeValue
	PartySize []interface{} `json:"partySize,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to [[PriceSpecification]] and its subtypes.
	// types : Text
	PriceCurrency []string `json:"priceCurrency,omitempty"`

	// ProgramMembershipUsed see : https://schema.org/programMembershipUsed
	// Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
	// types : ProgramMembership
	ProgramMembershipUsed []*ProgramMembership `json:"programMembershipUsed,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	// types : Person Organization
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

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// StartTime see : https://schema.org/startTime
	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from *January* to December.\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	StartTime []DateTime `json:"startTime,omitempty"`

	// TotalPrice see : https://schema.org/totalPrice
	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	// types : Number Text PriceSpecification
	TotalPrice []interface{} `json:"totalPrice,omitempty"`

	// UnderName see : https://schema.org/underName
	// The person or organization the reservation or ticket is for.
	// types : Person Organization
	UnderName []interface{} `json:"underName,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v FoodEstablishmentReservation) MarshalJSON() ([]byte, error) {
	type Alias FoodEstablishmentReservation

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"FoodEstablishmentReservation\","), b[1:]...), nil
}
