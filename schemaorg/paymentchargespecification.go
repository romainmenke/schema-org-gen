package schemaorg

import "encoding/json"

// PaymentChargeSpecification see : https://schema.org/PaymentChargeSpecification
type PaymentChargeSpecification struct {

typeContext

PriceSpecification

// AppliesToDeliveryMethod see : https://schema.org/appliesToDeliveryMethod
// The delivery method(s) to which the delivery charge or payment charge specification applies.
AppliesToDeliveryMethod *DeliveryMethod `json:"appliesToDeliveryMethod"`

// AppliesToPaymentMethod see : https://schema.org/appliesToPaymentMethod
// The payment method(s) to which the payment charge specification applies.
AppliesToPaymentMethod *PaymentMethod `json:"appliesToPaymentMethod"`

}

func (v *PaymentChargeSpecification) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentChargeSpecification"

	return json.Marshal(v)
}
