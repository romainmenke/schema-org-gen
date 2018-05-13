package schemaorggo

import "encoding/json"

// ReceiveAction see : https://schema.org/ReceiveAction
type ReceiveAction struct {
	TransferAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	// types : DeliveryMethod
	DeliveryMethod []*DeliveryMethod `json:"deliveryMethod,omitempty"`

	// Sender see : https://schema.org/sender
	// A sub property of participant. The participant who is at the sending end of the action.
	// types : Audience Organization Person
	Sender []interface{} `json:"sender,omitempty"`
}

func (v ReceiveAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TransferAction.intoMap(intop)

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

	{
		var value interface{} = v.Sender
		if len(v.Sender) == 1 {
			value = v.Sender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sender"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReceiveAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReceiveAction"

	return data, nil
}

func (v ReceiveAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
