package schemaorggo

import "encoding/json"

// UnitPriceSpecification see : https://schema.org/UnitPriceSpecification
type UnitPriceSpecification struct {
	PriceSpecification

	typeContext

	// BillingIncrement see : https://schema.org/billingIncrement
	// This property specifies the minimal quantity and rounding increment that will be the basis for the billing. The unit of measurement is specified by the unitCode property.
	// types : Number
	BillingIncrement []float64 `json:"billingIncrement,omitempty"`

	// PriceType see : https://schema.org/priceType
	// A short text or acronym indicating multiple price specifications for the same offer, e.g. SRP for the suggested retail price or INVOICE for the invoice price, mostly used in the car industry.
	// types : Text
	PriceType []string `json:"priceType,omitempty"`

	// ReferenceQuantity see : https://schema.org/referenceQuantity
	// The reference quantity for which a certain price applies, e.g. 1 EUR per 4 kWh of electricity. This property is a replacement for unitOfMeasurement for the advanced cases where the price does not relate to a standard unit.
	// types : QuantitativeValue
	ReferenceQuantity []*QuantitativeValue `json:"referenceQuantity,omitempty"`

	// UnitCode see : https://schema.org/unitCode
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	// types : Text URL
	UnitCode []string `json:"unitCode,omitempty"`

	// UnitText see : https://schema.org/unitText
	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	// unitCode (see: https://schema.orgunitCode).
	// types : Text
	UnitText []string `json:"unitText,omitempty"`
}

func (v UnitPriceSpecification) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PriceSpecification.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.BillingIncrement
		if len(v.BillingIncrement) == 1 {
			value = v.BillingIncrement[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["billingIncrement"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PriceType
		if len(v.PriceType) == 1 {
			value = v.PriceType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReferenceQuantity
		if len(v.ReferenceQuantity) == 1 {
			value = v.ReferenceQuantity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["referenceQuantity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UnitCode
		if len(v.UnitCode) == 1 {
			value = v.UnitCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unitCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UnitText
		if len(v.UnitText) == 1 {
			value = v.UnitText[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unitText"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v UnitPriceSpecification) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "UnitPriceSpecification"

	return data, nil
}

func (v UnitPriceSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
