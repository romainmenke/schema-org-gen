package schemaorggo

import "encoding/json"

// DeliveryChargeSpecification see : https://schema.org/DeliveryChargeSpecification
type DeliveryChargeSpecification struct {
	PriceSpecification

	typeContext

	// AppliesToDeliveryMethod see : https://schema.org/appliesToDeliveryMethod
	// The delivery method(s) to which the delivery charge or payment charge specification applies.
	// types : DeliveryMethod
	AppliesToDeliveryMethod []*DeliveryMethod `json:"appliesToDeliveryMethod,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	// types : AdministrativeArea GeoShape Place Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// EligibleRegion see : https://pending.schema.org/eligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
	//
	// See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
	// types : GeoShape Place Text
	EligibleRegion []interface{} `json:"eligibleRegion,omitempty"`

	// IneligibleRegion see : https://schema.org/ineligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
	//
	// See also eligibleRegion (see: https://schema.org/eligibleRegion).
	// types : GeoShape Place Text
	IneligibleRegion []interface{} `json:"ineligibleRegion,omitempty"`
}

func (v DeliveryChargeSpecification) intoMap(intop *map[string]interface{}) error {
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
		var value interface{} = v.AreaServed
		if len(v.AreaServed) == 1 {
			value = v.AreaServed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["areaServed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EligibleRegion
		if len(v.EligibleRegion) == 1 {
			value = v.EligibleRegion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eligibleRegion"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IneligibleRegion
		if len(v.IneligibleRegion) == 1 {
			value = v.IneligibleRegion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ineligibleRegion"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v DeliveryChargeSpecification) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DeliveryChargeSpecification"

	return data, nil
}

func (v DeliveryChargeSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
