package schemaorggo

import "encoding/json"

// WarrantyPromise see : https://schema.org/WarrantyPromise
type WarrantyPromise struct {
	StructuredValue

	typeContext

	// DurationOfWarranty see : https://schema.org/durationOfWarranty
	// The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days.
	// types : QuantitativeValue
	DurationOfWarranty []*QuantitativeValue `json:"durationOfWarranty,omitempty"`

	// WarrantyScope see : https://schema.org/warrantyScope
	// The scope of the warranty promise.
	// types : WarrantyScope
	WarrantyScope []*WarrantyScope `json:"warrantyScope,omitempty"`
}

func (v WarrantyPromise) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.DurationOfWarranty
		if len(v.DurationOfWarranty) == 1 {
			value = v.DurationOfWarranty[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["durationOfWarranty"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WarrantyScope
		if len(v.WarrantyScope) == 1 {
			value = v.WarrantyScope[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["warrantyScope"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v WarrantyPromise) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WarrantyPromise"

	return data, nil
}

func (v WarrantyPromise) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
