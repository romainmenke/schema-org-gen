package schemaorggo

import "encoding/json"

// HowToItem see : https://schema.org/HowToItem
type HowToItem struct {
	ListItem

	typeContext

	// RequiredQuantity see : https://schema.org/requiredQuantity
	// The required quantity of the item(s).
	// types : Number QuantitativeValue Text
	RequiredQuantity []interface{} `json:"requiredQuantity,omitempty"`
}

func (v HowToItem) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ListItem.IntoMap(intop)

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

func (v HowToItem) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowToItem"

	return data, nil
}

func (v HowToItem) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
