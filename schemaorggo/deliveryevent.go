package schemaorggo

import "encoding/json"

// DeliveryEvent see : https://schema.org/DeliveryEvent
type DeliveryEvent struct {
	Event

	typeContext

	// AccessCode see : https://schema.org/accessCode
	// Password, PIN, or access code needed for delivery (e.g. from a locker).
	// types : Text
	AccessCode []string `json:"accessCode,omitempty"`

	// AvailableFrom see : https://schema.org/availableFrom
	// When the item is available for pickup from the store, locker, etc.
	// types : DateTime
	AvailableFrom []DateTime `json:"availableFrom,omitempty"`

	// AvailableThrough see : https://schema.org/availableThrough
	// After this date, the item will no longer be available for pickup.
	// types : DateTime
	AvailableThrough []DateTime `json:"availableThrough,omitempty"`

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	// types : DeliveryMethod
	HasDeliveryMethod []*DeliveryMethod `json:"hasDeliveryMethod,omitempty"`
}

func (v DeliveryEvent) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Event.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AccessCode
		if len(v.AccessCode) == 1 {
			value = v.AccessCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailableFrom
		if len(v.AvailableFrom) == 1 {
			value = v.AvailableFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailableThrough
		if len(v.AvailableThrough) == 1 {
			value = v.AvailableThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableThrough"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasDeliveryMethod
		if len(v.HasDeliveryMethod) == 1 {
			value = v.HasDeliveryMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasDeliveryMethod"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v DeliveryEvent) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DeliveryEvent"

	return data, nil
}

func (v DeliveryEvent) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
