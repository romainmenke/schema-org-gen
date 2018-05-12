package schemaorggo

import "encoding/json"

// TrainTrip see : https://schema.org/TrainTrip
type TrainTrip struct {
	Intangible

	typeContext

	// ArrivalPlatform see : https://schema.org/arrivalPlatform
	// The platform where the train arrives.
	ArrivalPlatform string `json:"arrivalPlatform"`

	// ArrivalStation see : https://schema.org/arrivalStation
	// The station where the train trip ends.
	ArrivalStation *TrainStation `json:"arrivalStation"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	ArrivalTime DateTime `json:"arrivalTime"`

	// DeparturePlatform see : https://schema.org/departurePlatform
	// The platform from which the train departs.
	DeparturePlatform string `json:"departurePlatform"`

	// DepartureStation see : https://schema.org/departureStation
	// The station from which the train departs.
	DepartureStation *TrainStation `json:"departureStation"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	DepartureTime DateTime `json:"departureTime"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	Provider interface{} `json:"provider"` // types : Organization Person

	// TrainName see : https://schema.org/trainName
	// The name of the train (e.g. The Orient Express).
	TrainName string `json:"trainName"`

	// TrainNumber see : https://schema.org/trainNumber
	// The unique identifier for the train.
	TrainNumber string `json:"trainNumber"`
}

func (v *TrainTrip) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TrainTrip"

	return json.Marshal(v)
}
