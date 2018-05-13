package schemaorggo

import "encoding/json"

// HowToTool see : https://schema.org/HowToTool
type HowToTool struct {
	HowToItem

	typeContext

	// RequiredQuantity see : https://schema.org/requiredQuantity
	// The required quantity of the item(s).
	// types : Number QuantitativeValue Text
	RequiredQuantity []interface{} `json:"requiredQuantity,omitempty"`
}

func (v HowToTool) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.HowToItem.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.RequiredQuantity
		if len(v.RequiredQuantity) == 1 {
			value = v.RequiredQuantity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredQuantity"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v HowToTool) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowToTool"

	return data, nil
}

func (v HowToTool) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
