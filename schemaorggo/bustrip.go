package schemaorggo

import "encoding/json"

// BusTrip see : https://schema.org/BusTrip
type BusTrip struct {
	Intangible

	typeContext

	// ArrivalBusStop see : https://schema.org/arrivalBusStop
	// The stop or station from which the bus arrives.
	// types : BusStation BusStop
	ArrivalBusStop interface{} `json:"arrivalBusStop,omitempty"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	// types : DateTime
	ArrivalTime DateTime `json:"arrivalTime,omitempty"`

	// BusName see : https://schema.org/busName
	// The name of the bus (e.g. Bolt Express).
	// types : Text
	BusName string `json:"busName,omitempty"`

	// BusNumber see : https://schema.org/busNumber
	// The unique identifier for the bus.
	// types : Text
	BusNumber string `json:"busNumber,omitempty"`

	// DepartureBusStop see : https://schema.org/departureBusStop
	// The stop or station from which the bus departs.
	// types : BusStation BusStop
	DepartureBusStop interface{} `json:"departureBusStop,omitempty"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	// types : DateTime
	DepartureTime DateTime `json:"departureTime,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider interface{} `json:"provider,omitempty"`
}

func (v BusTrip) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusTrip"

	return json.Marshal(v)
}

func (v *BusTrip) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BusTrip"

	return json.Marshal(*v)
}
