package schemaorg

import "encoding/json"

// DeliveryEvent see : https://schema.org/DeliveryEvent
type DeliveryEvent struct {

typeContext

Event

// AccessCode see : /accessCode
// Password, PIN, or access code needed for delivery (e.g. from a locker).
AccessCode string `json:"accessCode"`

// AvailableFrom see : /availableFrom
// When the item is available for pickup from the store, locker, etc.
AvailableFrom interface{} `json:"availableFrom"`

// AvailableThrough see : /availableThrough
// After this date, the item will no longer be available for pickup.
AvailableThrough interface{} `json:"availableThrough"`

// HasDeliveryMethod see : /hasDeliveryMethod
// Method used for delivery or shipping.
HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod"`

}

func (v *DeliveryEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeliveryEvent"

	return json.Marshal(v)
}
