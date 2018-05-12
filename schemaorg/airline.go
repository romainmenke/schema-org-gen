package schemaorg

import "encoding/json"

// Airline see : https://schema.org/Airline
type Airline struct {

Organization

typeContext

// BoardingPolicy see : https://schema.org/boardingPolicy
// The type of boarding policy used by the airline (e.g. zone-based or group-based).
BoardingPolicy *BoardingPolicyType `json:"boardingPolicy"`

// IataCode see : https://schema.org/iataCode
// IATA identifier for an airline or airport.
IataCode string `json:"iataCode"`

}

func (v *Airline) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Airline"

	return json.Marshal(v)
}
