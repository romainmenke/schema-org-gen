package schemaorggo

import "encoding/json"

// PaymentChargeSpecification see : https://schema.org/PaymentChargeSpecification
type PaymentChargeSpecification struct {
	PriceSpecification

	typeContext

	// AppliesToDeliveryMethod see : https://schema.org/appliesToDeliveryMethod
	// The delivery method(s) to which the delivery charge or payment charge specification applies.
	// types : DeliveryMethod
	AppliesToDeliveryMethod []*DeliveryMethod `json:"appliesToDeliveryMethod,omitempty"`

	// AppliesToPaymentMethod see : https://schema.org/appliesToPaymentMethod
	// The payment method(s) to which the payment charge specification applies.
	// types : PaymentMethod
	AppliesToPaymentMethod []*PaymentMethod `json:"appliesToPaymentMethod,omitempty"`
}

func (v PaymentChargeSpecification) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PriceSpecification.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AppliesToDeliveryMethod
		if len(v.AppliesToDeliveryMethod) == 1 {
			value = v.AppliesToDeliveryMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["appliesToDeliveryMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AppliesToPaymentMethod
		if len(v.AppliesToPaymentMethod) == 1 {
			value = v.AppliesToPaymentMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["appliesToPaymentMethod"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PaymentChargeSpecification) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PaymentChargeSpecification"

	return data, nil
}

func (v PaymentChargeSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
