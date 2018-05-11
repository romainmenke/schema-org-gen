package schemaorg

// Flight see : https://schema.org/Flight
type Flight struct {

Intangible// Aircraft see : /aircraft
// The kind of aircraft (e.g., "Boeing 747").
Aircraft interface{} `json:"aircraft"` // types : Text Vehicle

// ArrivalAirport see : /arrivalAirport
// The airport where the flight terminates.
ArrivalAirport string `json:"arrivalAirport"`

// ArrivalGate see : /arrivalGate
// Identifier of the flight's arrival gate.
ArrivalGate string `json:"arrivalGate"`

// ArrivalTerminal see : /arrivalTerminal
// Identifier of the flight's arrival terminal.
ArrivalTerminal string `json:"arrivalTerminal"`

// ArrivalTime see : /arrivalTime
// The expected arrival time.
ArrivalTime string `json:"arrivalTime"`

// BoardingPolicy see : /boardingPolicy
// The type of boarding policy used by the airline (e.g. zone-based or group-based).
BoardingPolicy string `json:"boardingPolicy"`

// DepartureAirport see : /departureAirport
// The airport where the flight originates.
DepartureAirport string `json:"departureAirport"`

// DepartureGate see : /departureGate
// Identifier of the flight's departure gate.
DepartureGate string `json:"departureGate"`

// DepartureTerminal see : /departureTerminal
// Identifier of the flight's departure terminal.
DepartureTerminal string `json:"departureTerminal"`

// DepartureTime see : /departureTime
// The expected departure time.
DepartureTime string `json:"departureTime"`

// EstimatedFlightDuration see : /estimatedFlightDuration
// The estimated time the flight will take.
EstimatedFlightDuration interface{} `json:"estimatedFlightDuration"` // types : Duration Text

// FlightDistance see : /flightDistance
// The distance of the flight.
FlightDistance interface{} `json:"flightDistance"` // types : Distance Text

// FlightNumber see : /flightNumber
// The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is 'UA', the flightNumber is 'UA110'.
FlightNumber string `json:"flightNumber"`

// MealService see : /mealService
// Description of the meals that will be provided or available for purchase.
MealService string `json:"mealService"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// Seller see : /seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
Seller interface{} `json:"seller"` // types : Organization Person

// WebCheckinTime see : /webCheckinTime
// The time when a passenger can check into the flight online.
WebCheckinTime string `json:"webCheckinTime"`

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

}

