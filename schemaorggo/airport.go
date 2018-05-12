package schemaorggo

import "encoding/json"

// Airport see : https://schema.org/Airport
type Airport struct {
	CivicStructure

	typeContext

	// IataCode see : https://schema.org/iataCode
	// IATA identifier for an airline or airport.
	IataCode string `json:"iataCode"`

	// IcaoCode see : https://schema.org/icaoCode
	// ICAO identifier for an airport.
	IcaoCode string `json:"icaoCode"`
}

func (v *Airport) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Airport"

	return json.Marshal(v)
}
