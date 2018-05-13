package schemaorggo

import "encoding/json"

// SellAction see : https://schema.org/SellAction
type SellAction struct {
	TradeAction

	typeContext

	// Buyer see : https://schema.org/buyer
	// A sub property of participant. The participant/person/organization that bought the object.
	// types : Person
	Buyer []*Person `json:"buyer,omitempty"`
}

func (v SellAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TradeAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Buyer
		if len(v.Buyer) == 1 {
			value = v.Buyer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["buyer"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SellAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SellAction"

	return data, nil
}

func (v SellAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
