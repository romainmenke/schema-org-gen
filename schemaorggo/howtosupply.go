package schemaorggo

import "encoding/json"

// HowToSupply see : https://schema.org/HowToSupply
type HowToSupply struct {
	HowToItem

	typeContext

	// EstimatedCost see : https://schema.org/estimatedCost
	// The estimated cost of the supply or supplies consumed when performing instructions.
	// types : MonetaryAmount Text
	EstimatedCost []interface{} `json:"estimatedCost,omitempty"`
}

func (v HowToSupply) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.HowToItem.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.EstimatedCost
		if len(v.EstimatedCost) == 1 {
			value = v.EstimatedCost[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["estimatedCost"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v HowToSupply) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowToSupply"

	return data, nil
}

func (v HowToSupply) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
