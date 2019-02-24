package schemaorggo

import "encoding/json"

// RentalCarReservation see : https://schema.org/RentalCarReservation
type RentalCarReservation struct {

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
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DropoffLocation see : https://schema.org/dropoffLocation
	// Where a rental car can be dropped off.
	// types : Place
	DropoffLocation []*Place `json:"dropoffLocation,omitempty"`

	// DropoffTime see : https://schema.org/dropoffTime
	// When a rental car can be dropped off.
	// types : DateTime
	DropoffTime []DateTime `json:"dropoffTime,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
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

	// PickupLocation see : https://schema.org/pickupLocation
	// Where a taxi will pick up a passenger or a rental car can be picked up.
	// types : Place
	PickupLocation []*Place `json:"pickupLocation,omitempty"`

	// PickupTime see : https://schema.org/pickupTime
	// When a taxi will pickup a passenger or a rental car can be picked up.
	// types : DateTime
	PickupTime []DateTime `json:"pickupTime,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification and its subtypes.
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

func (v RentalCarReservation) MarshalJSON() ([]byte, error) {
	type Alias RentalCarReservation

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"RentalCarReservation\","), b[1:]...), nil
}
