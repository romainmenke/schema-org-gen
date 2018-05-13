package schemaorggo

import "encoding/json"

// OrderAction see : https://schema.org/OrderAction
type OrderAction struct {
	TradeAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	// types : DeliveryMethod
	DeliveryMethod []*DeliveryMethod `json:"deliveryMethod,omitempty"`
}

func (v OrderAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TradeAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.DeliveryMethod
		if len(v.DeliveryMethod) == 1 {
			value = v.DeliveryMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["deliveryMethod"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v OrderAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "OrderAction"

	return data, nil
}

func (v OrderAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
