package schemaorg

import "encoding/json"

// TrainTrip see : https://schema.org/TrainTrip
type TrainTrip struct {

typeContext

Intangible

// ArrivalPlatform see : /arrivalPlatform
// The platform where the train arrives.
ArrivalPlatform string `json:"arrivalPlatform"`

// ArrivalStation see : /arrivalStation
// The station where the train trip ends.
ArrivalStation *TrainStation `json:"arrivalStation"`

// ArrivalTime see : /arrivalTime
// The expected arrival time.
ArrivalTime interface{} `json:"arrivalTime"`

// DeparturePlatform see : /departurePlatform
// The platform from which the train departs.
DeparturePlatform string `json:"departurePlatform"`

// DepartureStation see : /departureStation
// The station from which the train departs.
DepartureStation *TrainStation `json:"departureStation"`

// DepartureTime see : /departureTime
// The expected departure time.
DepartureTime interface{} `json:"departureTime"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// TrainName see : /trainName
// The name of the train (e.g. The Orient Express).
TrainName string `json:"trainName"`

// TrainNumber see : /trainNumber
// The unique identifier for the train.
TrainNumber string `json:"trainNumber"`

}

func (v *TrainTrip) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TrainTrip"

	return json.Marshal(v)
}
