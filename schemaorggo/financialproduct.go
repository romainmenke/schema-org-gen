package schemaorggo

import "encoding/json"

// FinancialProduct see : https://schema.org/FinancialProduct
type FinancialProduct struct {
	Service

	typeContext

	// AnnualPercentageRate see : https://schema.org/annualPercentageRate
	// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
	// types : Number QuantitativeValue
	AnnualPercentageRate []interface{} `json:"annualPercentageRate,omitempty"`

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	// types : Text URL
	FeesAndCommissionsSpecification []string `json:"feesAndCommissionsSpecification,omitempty"`

	// InterestRate see : https://schema.org/interestRate
	// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
	// types : Number QuantitativeValue
	InterestRate []interface{} `json:"interestRate,omitempty"`
}

func (v FinancialProduct) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Service.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AnnualPercentageRate
		if len(v.AnnualPercentageRate) == 1 {
			value = v.AnnualPercentageRate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["annualPercentageRate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FeesAndCommissionsSpecification
		if len(v.FeesAndCommissionsSpecification) == 1 {
			value = v.FeesAndCommissionsSpecification[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["feesAndCommissionsSpecification"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InterestRate
		if len(v.InterestRate) == 1 {
			value = v.InterestRate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["interestRate"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v FinancialProduct) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "FinancialProduct"

	return data, nil
}

func (v FinancialProduct) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
