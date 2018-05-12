package schemaorggo

import "encoding/json"

// RentAction see : https://schema.org/RentAction
type RentAction struct {
	TradeAction

	typeContext

	// Landlord see : https://schema.org/landlord
	// A sub property of participant. The owner of the real estate property.
	Landlord interface{} `json:"landlord,omitempty"` // types : Organization Person

	// RealEstateAgent see : https://schema.org/realEstateAgent
	// A sub property of participant. The real estate agent involved in the action.
	RealEstateAgent *RealEstateAgent `json:"realEstateAgent,omitempty"` // types : RealEstateAgent

}

func (v RentAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RentAction"

	return json.Marshal(v)
}

func (v *RentAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RentAction"

	return json.Marshal(*v)
}
