package schemaorggo

import "encoding/json"

// Airline see : https://schema.org/Airline
type Airline struct {
	Organization

	typeContext

	// BoardingPolicy see : https://schema.org/boardingPolicy
	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	// types : BoardingPolicyType
	BoardingPolicy *BoardingPolicyType `json:"boardingPolicy,omitempty"`

	// IataCode see : https://schema.org/iataCode
	// IATA identifier for an airline or airport.
	// types : Text
	IataCode string `json:"iataCode,omitempty"`
}

func (v Airline) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Airline"

	return json.Marshal(v)
}

func (v *Airline) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Airline"

	return json.Marshal(*v)
}
