package schemaorg

import "encoding/json"

// BusinessAudience see : https://schema.org/BusinessAudience
type BusinessAudience struct {

typeContext

Audience

// NumberOfEmployees see : https://schema.org/numberOfEmployees
// The number of employees in an organization e.g. business.
NumberOfEmployees *QuantitativeValue `json:"numberOfEmployees"`

// YearlyRevenue see : https://schema.org/yearlyRevenue
// The size of the business in annual revenue.
YearlyRevenue *QuantitativeValue `json:"yearlyRevenue"`

// YearsInOperation see : https://schema.org/yearsInOperation
// The age of the business.
YearsInOperation *QuantitativeValue `json:"yearsInOperation"`

}

func (v *BusinessAudience) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusinessAudience"

	return json.Marshal(v)
}
