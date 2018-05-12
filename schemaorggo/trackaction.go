package schemaorggo

import "encoding/json"

// TrackAction see : https://schema.org/TrackAction
type TrackAction struct {
	FindAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	// types : DeliveryMethod
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty"`
}

func (v TrackAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TrackAction"

	return json.Marshal(v)
}

func (v *TrackAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TrackAction"

	return json.Marshal(*v)
}
