package schemaorg

import "encoding/json"

// FoodEstablishmentReservation see : https://schema.org/FoodEstablishmentReservation
type FoodEstablishmentReservation struct {

typeContext

Reservation

// EndTime see : /endTime
// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
// 
// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
EndTime interface{} `json:"endTime"`

// PartySize see : /partySize
// Number of people the reservation should accommodate.
PartySize interface{} `json:"partySize"` // types : Integer QuantitativeValue

// StartTime see : /startTime
// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
// 
// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
StartTime interface{} `json:"startTime"`

}

func (v *FoodEstablishmentReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FoodEstablishmentReservation"

	return json.Marshal(v)
}
