package schemaorggo

import "encoding/json"

// BusTrip see : https://schema.org/BusTrip
type BusTrip struct {
	Intangible

	typeContext

	// ArrivalBusStop see : https://schema.org/arrivalBusStop
	// The stop or station from which the bus arrives.
	ArrivalBusStop interface{} `json:"arrivalBusStop,omitempty"` // types : BusStation BusStop

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	ArrivalTime DateTime `json:"arrivalTime,omitempty"` // types : DateTime

	// BusName see : https://schema.org/busName
	// The name of the bus (e.g. Bolt Express).
	BusName string `json:"busName,omitempty"` // types : Text

	// BusNumber see : https://schema.org/busNumber
	// The unique identifier for the bus.
	BusNumber string `json:"busNumber,omitempty"` // types : Text

	// DepartureBusStop see : https://schema.org/departureBusStop
	// The stop or station from which the bus departs.
	DepartureBusStop interface{} `json:"departureBusStop,omitempty"` // types : BusStation BusStop

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	DepartureTime DateTime `json:"departureTime,omitempty"` // types : DateTime

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	Provider interface{} `json:"provider,omitempty"` // types : Organization Person

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
