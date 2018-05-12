package schemaorggo

import "encoding/json"

// DeliveryEvent see : https://schema.org/DeliveryEvent
type DeliveryEvent struct {
	Event

	typeContext

	// AccessCode see : https://schema.org/accessCode
	// Password, PIN, or access code needed for delivery (e.g. from a locker).
	AccessCode string `json:"accessCode,omitempty"` // types : Text

	// AvailableFrom see : https://schema.org/availableFrom
	// When the item is available for pickup from the store, locker, etc.
	AvailableFrom DateTime `json:"availableFrom,omitempty"` // types : DateTime

	// AvailableThrough see : https://schema.org/availableThrough
	// After this date, the item will no longer be available for pickup.
	AvailableThrough DateTime `json:"availableThrough,omitempty"` // types : DateTime

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod,omitempty"` // types : DeliveryMethod

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
