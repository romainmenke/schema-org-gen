package schemaorggo

import "encoding/json"

// Flight see : https://schema.org/Flight
type Flight struct {
	Intangible

	typeContext

	// Aircraft see : https://schema.org/aircraft
	// The kind of aircraft (e.g., "Boeing 747").
	Aircraft interface{} `json:"aircraft,omitempty"` // types : Text Vehicle

	// ArrivalAirport see : https://schema.org/arrivalAirport
	// The airport where the flight terminates.
	ArrivalAirport *Airport `json:"arrivalAirport,omitempty"`

	// ArrivalGate see : https://schema.org/arrivalGate
	// Identifier of the flight's arrival gate.
	ArrivalGate string `json:"arrivalGate,omitempty"`

	// ArrivalTerminal see : https://schema.org/arrivalTerminal
	// Identifier of the flight's arrival terminal.
	ArrivalTerminal string `json:"arrivalTerminal,omitempty"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	ArrivalTime DateTime `json:"arrivalTime,omitempty"`

	// BoardingPolicy see : https://schema.org/boardingPolicy
	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	BoardingPolicy *BoardingPolicyType `json:"boardingPolicy,omitempty"`

	// DepartureAirport see : https://schema.org/departureAirport
	// The airport where the flight originates.
	DepartureAirport *Airport `json:"departureAirport,omitempty"`

	// DepartureGate see : https://schema.org/departureGate
	// Identifier of the flight's departure gate.
	DepartureGate string `json:"departureGate,omitempty"`

	// DepartureTerminal see : https://schema.org/departureTerminal
	// Identifier of the flight's departure terminal.
	DepartureTerminal string `json:"departureTerminal,omitempty"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	DepartureTime DateTime `json:"departureTime,omitempty"`

	// EstimatedFlightDuration see : https://schema.org/estimatedFlightDuration
	// The estimated time the flight will take.
	EstimatedFlightDuration interface{} `json:"estimatedFlightDuration,omitempty"` // types : Duration Text

	// FlightDistance see : https://schema.org/flightDistance
	// The distance of the flight.
	FlightDistance interface{} `json:"flightDistance,omitempty"` // types : Distance Text

	// FlightNumber see : https://schema.org/flightNumber
	// The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is 'UA', the flightNumber is 'UA110'.
	FlightNumber string `json:"flightNumber,omitempty"`

	// MealService see : https://schema.org/mealService
	// Description of the meals that will be provided or available for purchase.
	MealService string `json:"mealService,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	Provider interface{} `json:"provider,omitempty"` // types : Organization Person

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	Seller interface{} `json:"seller,omitempty"` // types : Organization Person

	// WebCheckinTime see : https://schema.org/webCheckinTime
	// The time when a passenger can check into the flight online.
	WebCheckinTime DateTime `json:"webCheckinTime,omitempty"`
}

func (v Flight) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Flight"

	return json.Marshal(v)
}

func (v *Flight) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Flight"

	return json.Marshal(*v)
}
