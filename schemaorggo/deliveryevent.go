package schemaorggo

import "encoding/json"

// DeliveryEvent see : https://schema.org/DeliveryEvent
type DeliveryEvent struct {
	Event

	typeContext

	// AccessCode see : https://schema.org/accessCode
	// Password, PIN, or access code needed for delivery (e.g. from a locker).
	// types : Text
	AccessCode string `json:"accessCode,omitempty"`

	// AvailableFrom see : https://schema.org/availableFrom
	// When the item is available for pickup from the store, locker, etc.
	// types : DateTime
	AvailableFrom DateTime `json:"availableFrom,omitempty"`

	// AvailableThrough see : https://schema.org/availableThrough
	// After this date, the item will no longer be available for pickup.
	// types : DateTime
	AvailableThrough DateTime `json:"availableThrough,omitempty"`

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	// types : DeliveryMethod
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod,omitempty"`
}

func (v DeliveryEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeliveryEvent"

	return json.Marshal(v)
}

func (v *DeliveryEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DeliveryEvent"

	return json.Marshal(*v)
}
