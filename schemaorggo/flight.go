package schemaorggo

import "encoding/json"

// Flight see : https://schema.org/Flight
type Flight struct {
	Intangible

	typeContext

	// Aircraft see : https://schema.org/aircraft
	// The kind of aircraft (e.g., &quot;Boeing 747&quot;).
	// types : Text Vehicle
	Aircraft []interface{} `json:"aircraft,omitempty"`

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

	// EstimatedFlightDuration see : https://schema.org/estimatedFlightDuration
	// The estimated time the flight will take.
	// types : Duration Text
	EstimatedFlightDuration []interface{} `json:"estimatedFlightDuration,omitempty"`

	// FlightDistance see : https://schema.org/flightDistance
	// The distance of the flight.
	// types : Distance Text
	FlightDistance []interface{} `json:"flightDistance,omitempty"`

	// FlightNumber see : https://schema.org/flightNumber
	// The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is &#39;UA&#39;, the flightNumber is &#39;UA110&#39;.
	// types : Text
	FlightNumber []string `json:"flightNumber,omitempty"`

	// MealService see : https://schema.org/mealService
	// Description of the meals that will be provided or available for purchase.
	// types : Text
	MealService []string `json:"mealService,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller []interface{} `json:"seller,omitempty"`

	// WebCheckinTime see : https://schema.org/webCheckinTime
	// The time when a passenger can check into the flight online.
	// types : DateTime
	WebCheckinTime []DateTime `json:"webCheckinTime,omitempty"`
}

func (v Flight) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Aircraft
		if len(v.Aircraft) == 1 {
			value = v.Aircraft[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["aircraft"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalAirport
		if len(v.ArrivalAirport) == 1 {
			value = v.ArrivalAirport[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalAirport"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalGate
		if len(v.ArrivalGate) == 1 {
			value = v.ArrivalGate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalGate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalTerminal
		if len(v.ArrivalTerminal) == 1 {
			value = v.ArrivalTerminal[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalTerminal"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalTime
		if len(v.ArrivalTime) == 1 {
			value = v.ArrivalTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BoardingPolicy
		if len(v.BoardingPolicy) == 1 {
			value = v.BoardingPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["boardingPolicy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureAirport
		if len(v.DepartureAirport) == 1 {
			value = v.DepartureAirport[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureAirport"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureGate
		if len(v.DepartureGate) == 1 {
			value = v.DepartureGate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureGate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureTerminal
		if len(v.DepartureTerminal) == 1 {
			value = v.DepartureTerminal[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureTerminal"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureTime
		if len(v.DepartureTime) == 1 {
			value = v.DepartureTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EstimatedFlightDuration
		if len(v.EstimatedFlightDuration) == 1 {
			value = v.EstimatedFlightDuration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["estimatedFlightDuration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FlightDistance
		if len(v.FlightDistance) == 1 {
			value = v.FlightDistance[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["flightDistance"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FlightNumber
		if len(v.FlightNumber) == 1 {
			value = v.FlightNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["flightNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MealService
		if len(v.MealService) == 1 {
			value = v.MealService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mealService"] = json.RawMessage(b)
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
		var value interface{} = v.Seller
		if len(v.Seller) == 1 {
			value = v.Seller[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seller"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WebCheckinTime
		if len(v.WebCheckinTime) == 1 {
			value = v.WebCheckinTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["webCheckinTime"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Flight) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Flight"

	return data, nil
}

func (v Flight) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
