package schemaorg

import "encoding/json"

// TrackAction see : https://schema.org/TrackAction
type TrackAction struct {

typeContext

FindAction

// DeliveryMethod see : https://schema.org/deliveryMethod
// A sub property of instrument. The method of delivery.
DeliveryMethod *DeliveryMethod `json:"deliveryMethod"`

}

func (v *TrackAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TrackAction"

	return json.Marshal(v)
}
