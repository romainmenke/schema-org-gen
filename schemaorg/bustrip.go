package schemaorg

import "encoding/json"

// BusTrip see : https://schema.org/BusTrip
type BusTrip struct {

typeContext

Intangible

// ArrivalBusStop see : https://schema.org/arrivalBusStop
// The stop or station from which the bus arrives.
ArrivalBusStop interface{} `json:"arrivalBusStop"` // types : BusStation BusStop

// ArrivalTime see : https://schema.org/arrivalTime
// The expected arrival time.
ArrivalTime interface{} `json:"arrivalTime"`

// BusName see : https://schema.org/busName
// The name of the bus (e.g. Bolt Express).
BusName string `json:"busName"`

// BusNumber see : https://schema.org/busNumber
// The unique identifier for the bus.
BusNumber string `json:"busNumber"`

// DepartureBusStop see : https://schema.org/departureBusStop
// The stop or station from which the bus departs.
DepartureBusStop interface{} `json:"departureBusStop"` // types : BusStation BusStop

// DepartureTime see : https://schema.org/departureTime
// The expected departure time.
DepartureTime interface{} `json:"departureTime"`

// Provider see : https://schema.org/provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

}

func (v *BusTrip) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusTrip"

	return json.Marshal(v)
}
