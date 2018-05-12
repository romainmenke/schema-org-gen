package schemaorggo

import "encoding/json"

// Seat see : https://schema.org/Seat
type Seat struct {
	Intangible

	typeContext

	// SeatNumber see : https://schema.org/seatNumber
	// The location of the reserved seat (e.g., 27).
	SeatNumber string `json:"seatNumber,omitempty"` // types : Text

	// SeatRow see : https://schema.org/seatRow
	// The row location of the reserved seat (e.g., B).
	SeatRow string `json:"seatRow,omitempty"` // types : Text

	// SeatSection see : https://schema.org/seatSection
	// The section location of the reserved seat (e.g. Orchestra).
	SeatSection string `json:"seatSection,omitempty"` // types : Text

	// SeatingType see : https://schema.org/seatingType
	// The type/class of the seat.
	SeatingType interface{} `json:"seatingType,omitempty"` // types : QualitativeValue Text

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
