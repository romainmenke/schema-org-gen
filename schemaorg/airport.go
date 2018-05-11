package schemaorg

import "encoding/json"

// Airport see : https://schema.org/Airport
type Airport struct {

typeContext

CivicStructure

// IataCode see : /iataCode
// IATA identifier for an airline or airport.
IataCode string `json:"iataCode"`

// IcaoCode see : /icaoCode
// ICAO identifier for an airport.
IcaoCode string `json:"icaoCode"`

}

func (v *Airport) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Airport"

	return json.Marshal(v)
}
