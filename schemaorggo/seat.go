package schemaorggo

import "encoding/json"

// Seat see : https://schema.org/Seat
type Seat struct {
	Intangible

	typeContext

	// SeatNumber see : https://schema.org/seatNumber
	// The location of the reserved seat (e.g., 27).
	// types : Text
	SeatNumber string `json:"seatNumber,omitempty"`

	// SeatRow see : https://schema.org/seatRow
	// The row location of the reserved seat (e.g., B).
	// types : Text
	SeatRow string `json:"seatRow,omitempty"`

	// SeatSection see : https://schema.org/seatSection
	// The section location of the reserved seat (e.g. Orchestra).
	// types : Text
	SeatSection string `json:"seatSection,omitempty"`

	// SeatingType see : https://schema.org/seatingType
	// The type/class of the seat.
	// types : QualitativeValue Text
	SeatingType interface{} `json:"seatingType,omitempty"`
}

func (v Seat) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Seat"

	return json.Marshal(v)
}

func (v *Seat) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Seat"

	return json.Marshal(*v)
}
