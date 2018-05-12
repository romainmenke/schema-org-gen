package schemaorg

import "encoding/json"

// RentAction see : https://schema.org/RentAction
type RentAction struct {

TradeAction

typeContext

// Landlord see : https://schema.org/landlord
// A sub property of participant. The owner of the real estate property.
Landlord interface{} `json:"landlord"` // types : Organization Person

// RealEstateAgent see : https://schema.org/realEstateAgent
// A sub property of participant. The real estate agent involved in the action.
RealEstateAgent *RealEstateAgent `json:"realEstateAgent"`

}

func (v *RentAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RentAction"

	return json.Marshal(v)
}
