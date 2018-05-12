package schemaorggo

import "encoding/json"

// PaymentChargeSpecification see : https://schema.org/PaymentChargeSpecification
type PaymentChargeSpecification struct {
	PriceSpecification

	typeContext

	// AppliesToDeliveryMethod see : https://schema.org/appliesToDeliveryMethod
	// The delivery method(s) to which the delivery charge or payment charge specification applies.
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliesToDeliveryMethod,omitempty"`

	// AppliesToPaymentMethod see : https://schema.org/appliesToPaymentMethod
	// The payment method(s) to which the payment charge specification applies.
	AppliesToPaymentMethod *PaymentMethod `json:"appliesToPaymentMethod,omitempty"`
}

func (v PaymentChargeSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentChargeSpecification"

	return json.Marshal(v)
}

func (v *PaymentChargeSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PaymentChargeSpecification"

	return json.Marshal(*v)
}
