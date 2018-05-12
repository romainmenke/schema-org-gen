package schemaorggo

import "encoding/json"

// DeliveryEvent see : https://schema.org/DeliveryEvent
type DeliveryEvent struct {
	Event

	typeContext

	// AccessCode see : https://schema.org/accessCode
	// Password, PIN, or access code needed for delivery (e.g. from a locker).
	AccessCode string `json:"accessCode"`

	// AvailableFrom see : https://schema.org/availableFrom
	// When the item is available for pickup from the store, locker, etc.
	AvailableFrom interface{} `json:"availableFrom"`

	// AvailableThrough see : https://schema.org/availableThrough
	// After this date, the item will no longer be available for pickup.
	AvailableThrough interface{} `json:"availableThrough"`

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod"`
}

func (v *DeliveryEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeliveryEvent"

	return json.Marshal(v)
}
