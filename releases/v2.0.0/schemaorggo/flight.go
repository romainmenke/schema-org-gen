package schemaorggo

import "encoding/json"

// Flight see : https://schema.org/Flight
type Flight struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// Aircraft see : https://schema.org/aircraft
	// The kind of aircraft (e.g., &quot;Boeing 747&quot;).
	// types : Text Vehicle
	Aircraft []interface{} `json:"aircraft,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// ArrivalAirport see : https://schema.org/arrivalAirport
	// The airport where the flight terminates.
	// types : Airport
	ArrivalAirport []*Airport `json:"arrivalAirport,omitempty"`

	// ArrivalGate see : https://schema.org/arrivalGate
	// Identifier of the flight&#39;s arrival gate.
	// types : Text
	ArrivalGate []string `json:"arrivalGate,omitempty"`

	// ArrivalTerminal see : https://schema.org/arrivalTerminal
	// Identifier of the flight&#39;s arrival terminal.
	// types : Text
	ArrivalTerminal []string `json:"arrivalTerminal,omitempty"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	// types : DateTime
	ArrivalTime []DateTime `json:"arrivalTime,omitempty"`

	// BoardingPolicy see : https://schema.org/boardingPolicy
	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	// types : BoardingPolicyType
	BoardingPolicy []*BoardingPolicyType `json:"boardingPolicy,omitempty"`

	// Carrier see : https://schema.org/carrier
	// &#39;carrier&#39; is an out-dated term indicating the &#39;provider&#39; for parcel delivery and flights.
	// types : Organization
	Carrier []*Organization `json:"carrier,omitempty"`

	// DepartureAirport see : https://schema.org/departureAirport
	// The airport where the flight originates.
	// types : Airport
	DepartureAirport []*Airport `json:"departureAirport,omitempty"`

	// DepartureGate see : https://schema.org/departureGate
	// Identifier of the flight&#39;s departure gate.
	// types : Text
	DepartureGate []string `json:"departureGate,omitempty"`

	// DepartureTerminal see : https://schema.org/departureTerminal
	// Identifier of the flight&#39;s departure terminal.
	// types : Text
	DepartureTerminal []string `json:"departureTerminal,omitempty"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	// types : DateTime
	DepartureTime []DateTime `json:"departureTime,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// EstimatedFlightDuration see : https://schema.org/estimatedFlightDuration
	// The estimated time the flight will take.
	// types : Text Duration
	EstimatedFlightDuration []interface{} `json:"estimatedFlightDuration,omitempty"`

	// FlightDistance see : https://schema.org/flightDistance
	// The distance of the flight.
	// types : Text Distance
	FlightDistance []interface{} `json:"flightDistance,omitempty"`

	// FlightNumber see : https://schema.org/flightNumber
	// The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is &#39;UA&#39;, the flightNumber is &#39;UA110&#39;.
	// types : Text
	FlightNumber []string `json:"flightNumber,omitempty"`

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

	// MealService see : https://schema.org/mealService
	// Description of the meals that will be provided or available for purchase.
	// types : Text
	MealService []string `json:"mealService,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	// types : Person Organization
	Provider []interface{} `json:"provider,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	// types : Organization Person
	Seller []interface{} `json:"seller,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// WebCheckinTime see : https://schema.org/webCheckinTime
	// The time when a passenger can check into the flight online.
	// types : DateTime
	WebCheckinTime []DateTime `json:"webCheckinTime,omitempty"`
}

func (v Flight) MarshalJSON() ([]byte, error) {
	type Alias Flight

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Flight\","), b[1:]...), nil
}
