package schemaorg

import "encoding/json"

// Seat see : https://schema.org/Seat
type Seat struct {

typeContext

Intangible

// SeatNumber see : /seatNumber
// The location of the reserved seat (e.g., 27).
SeatNumber string `json:"seatNumber"`

// SeatRow see : /seatRow
// The row location of the reserved seat (e.g., B).
SeatRow string `json:"seatRow"`

// SeatSection see : /seatSection
// The section location of the reserved seat (e.g. Orchestra).
SeatSection string `json:"seatSection"`

// SeatingType see : /seatingType
// The type/class of the seat.
SeatingType interface{} `json:"seatingType"` // types : QualitativeValue Text

}

func (v *Seat) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Seat"

	return json.Marshal(v)
}
