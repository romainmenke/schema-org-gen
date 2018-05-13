package schemaorggo

import "encoding/json"

// FoodEstablishmentReservation see : https://schema.org/FoodEstablishmentReservation
type FoodEstablishmentReservation struct {
	Reservation

	typeContext

	// EndTime see : https://schema.org/endTime
	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	EndTime []DateTime `json:"endTime,omitempty"`

	// PartySize see : https://schema.org/partySize
	// Number of people the reservation should accommodate.
	// types : Integer QuantitativeValue
	PartySize []interface{} `json:"partySize,omitempty"`

	// StartTime see : https://schema.org/startTime
	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	StartTime []DateTime `json:"startTime,omitempty"`
}

func (v FoodEstablishmentReservation) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Reservation.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.EndTime
		if len(v.EndTime) == 1 {
			value = v.EndTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartySize
		if len(v.PartySize) == 1 {
			value = v.PartySize[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partySize"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StartTime
		if len(v.StartTime) == 1 {
			value = v.StartTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["startTime"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v FoodEstablishmentReservation) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "FoodEstablishmentReservation"

	return data, nil
}

func (v FoodEstablishmentReservation) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
