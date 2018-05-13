package schemaorggo

import "encoding/json"

// RentAction see : https://schema.org/RentAction
type RentAction struct {
	TradeAction

	typeContext

	// Landlord see : https://schema.org/landlord
	// A sub property of participant. The owner of the real estate property.
	// types : Organization Person
	Landlord []interface{} `json:"landlord,omitempty"`

	// RealEstateAgent see : https://schema.org/realEstateAgent
	// A sub property of participant. The real estate agent involved in the action.
	// types : RealEstateAgent
	RealEstateAgent []*RealEstateAgent `json:"realEstateAgent,omitempty"`
}

func (v RentAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TradeAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Landlord
		if len(v.Landlord) == 1 {
			value = v.Landlord[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["landlord"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RealEstateAgent
		if len(v.RealEstateAgent) == 1 {
			value = v.RealEstateAgent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["realEstateAgent"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v RentAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "RentAction"

	return data, nil
}

func (v RentAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
