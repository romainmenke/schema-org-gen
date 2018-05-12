package schemaorggo

import "encoding/json"

// BusinessAudience see : https://schema.org/BusinessAudience
type BusinessAudience struct {
	Audience

	typeContext

	// NumberOfEmployees see : https://schema.org/numberOfEmployees
	// The number of employees in an organization e.g. business.
	// types : QuantitativeValue
	NumberOfEmployees *QuantitativeValue `json:"numberOfEmployees,omitempty"`

	// YearlyRevenue see : https://schema.org/yearlyRevenue
	// The size of the business in annual revenue.
	// types : QuantitativeValue
	YearlyRevenue *QuantitativeValue `json:"yearlyRevenue,omitempty"`

	// YearsInOperation see : https://schema.org/yearsInOperation
	// The age of the business.
	// types : QuantitativeValue
	YearsInOperation *QuantitativeValue `json:"yearsInOperation,omitempty"`
}

func (v BusinessAudience) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusinessAudience"

	return json.Marshal(v)
}

func (v *BusinessAudience) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BusinessAudience"

	return json.Marshal(*v)
}
