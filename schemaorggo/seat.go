package schemaorggo

import "encoding/json"

// Seat see : https://schema.org/Seat
type Seat struct {
	Intangible

	typeContext

	// SeatNumber see : https://schema.org/seatNumber
	// The location of the reserved seat (e.g., 27).
	// types : Text
	SeatNumber []string `json:"seatNumber,omitempty"`

	// SeatRow see : https://schema.org/seatRow
	// The row location of the reserved seat (e.g., B).
	// types : Text
	SeatRow []string `json:"seatRow,omitempty"`

	// SeatSection see : https://schema.org/seatSection
	// The section location of the reserved seat (e.g. Orchestra).
	// types : Text
	SeatSection []string `json:"seatSection,omitempty"`

	// SeatingType see : https://schema.org/seatingType
	// The type/class of the seat.
	// types : QualitativeValue Text
	SeatingType []interface{} `json:"seatingType,omitempty"`
}

func (v Seat) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.SeatNumber
		if len(v.SeatNumber) == 1 {
			value = v.SeatNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seatNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SeatRow
		if len(v.SeatRow) == 1 {
			value = v.SeatRow[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seatRow"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SeatSection
		if len(v.SeatSection) == 1 {
			value = v.SeatSection[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seatSection"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SeatingType
		if len(v.SeatingType) == 1 {
			value = v.SeatingType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seatingType"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Seat) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Seat"

	return data, nil
}

func (v Seat) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
