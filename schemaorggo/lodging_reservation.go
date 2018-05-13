package schemaorggo

import "encoding/json"

// LodgingReservation see : https://schema.org/LodgingReservation
type LodgingReservation struct {
	Reservation

	typeContext

	// CheckinTime see : https://schema.org/checkinTime
	// The earliest someone may check into a lodging establishment.
	// types : DateTime
	CheckinTime []DateTime `json:"checkinTime,omitempty"`

	// CheckoutTime see : https://schema.org/checkoutTime
	// The latest someone may check out of a lodging establishment.
	// types : DateTime
	CheckoutTime []DateTime `json:"checkoutTime,omitempty"`

	// LodgingUnitDescription see : https://schema.org/lodgingUnitDescription
	// A full description of the lodging unit.
	// types : Text
	LodgingUnitDescription []string `json:"lodgingUnitDescription,omitempty"`

	// LodgingUnitType see : https://schema.org/lodgingUnitType
	// Textual description of the unit type (including suite vs. room, size of bed, etc.).
	// types : QualitativeValue Text
	LodgingUnitType []interface{} `json:"lodgingUnitType,omitempty"`

	// NumAdults see : https://schema.org/numAdults
	// The number of adults staying in the unit.
	// types : Integer QuantitativeValue
	NumAdults []interface{} `json:"numAdults,omitempty"`

	// NumChildren see : https://schema.org/numChildren
	// The number of children staying in the unit.
	// types : Integer QuantitativeValue
	NumChildren []interface{} `json:"numChildren,omitempty"`
}

func (v LodgingReservation) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Reservation.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CheckinTime
		if len(v.CheckinTime) == 1 {
			value = v.CheckinTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["checkinTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CheckoutTime
		if len(v.CheckoutTime) == 1 {
			value = v.CheckoutTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["checkoutTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LodgingUnitDescription
		if len(v.LodgingUnitDescription) == 1 {
			value = v.LodgingUnitDescription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lodgingUnitDescription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LodgingUnitType
		if len(v.LodgingUnitType) == 1 {
			value = v.LodgingUnitType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lodgingUnitType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumAdults
		if len(v.NumAdults) == 1 {
			value = v.NumAdults[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numAdults"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumChildren
		if len(v.NumChildren) == 1 {
			value = v.NumChildren[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numChildren"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v LodgingReservation) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "LodgingReservation"

	return data, nil
}

func (v LodgingReservation) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
