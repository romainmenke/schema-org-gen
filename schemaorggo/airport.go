package schemaorggo

import "encoding/json"

// Airport see : https://schema.org/Airport
type Airport struct {
	CivicStructure

	typeContext

	// IataCode see : https://schema.org/iataCode
	// IATA identifier for an airline or airport.
	// types : Text
	IataCode string `json:"iataCode,omitempty"`

	// IcaoCode see : https://schema.org/icaoCode
	// ICAO identifier for an airport.
	// types : Text
	IcaoCode string `json:"icaoCode,omitempty"`
}

func (v Airport) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Airport"

	return json.Marshal(v)
}

func (v *Airport) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Airport"

	return json.Marshal(*v)
}
