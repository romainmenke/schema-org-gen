package schemaorggo

import "encoding/json"

// BusinessAudience see : https://schema.org/BusinessAudience
type BusinessAudience struct {
	Audience

	typeContext

	// NumberOfEmployees see : https://schema.org/numberOfEmployees
	// The number of employees in an organization e.g. business.
	// types : QuantitativeValue
	NumberOfEmployees []*QuantitativeValue `json:"numberOfEmployees,omitempty"`

	// YearlyRevenue see : https://schema.org/yearlyRevenue
	// The size of the business in annual revenue.
	// types : QuantitativeValue
	YearlyRevenue []*QuantitativeValue `json:"yearlyRevenue,omitempty"`

	// YearsInOperation see : https://schema.org/yearsInOperation
	// The age of the business.
	// types : QuantitativeValue
	YearsInOperation []*QuantitativeValue `json:"yearsInOperation,omitempty"`
}

func (v BusinessAudience) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Audience.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.NumberOfEmployees
		if len(v.NumberOfEmployees) == 1 {
			value = v.NumberOfEmployees[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfEmployees"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.YearlyRevenue
		if len(v.YearlyRevenue) == 1 {
			value = v.YearlyRevenue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["yearlyRevenue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.YearsInOperation
		if len(v.YearsInOperation) == 1 {
			value = v.YearsInOperation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["yearsInOperation"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BusinessAudience) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BusinessAudience"

	return data, nil
}

func (v BusinessAudience) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
