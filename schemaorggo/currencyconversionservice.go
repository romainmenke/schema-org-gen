package schemaorggo

import "encoding/json"

// CurrencyConversionService see : https://schema.org/CurrencyConversionService
type CurrencyConversionService struct {
	FinancialProduct

	typeContext

	// AnnualPercentageRate see : https://schema.org/annualPercentageRate
	// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
	AnnualPercentageRate interface{} `json:"annualPercentageRate,omitempty"` // types : Number QuantitativeValue

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification,omitempty"` // types : Text URL

	// InterestRate see : https://schema.org/interestRate
	// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
	InterestRate interface{} `json:"interestRate,omitempty"` // types : Number QuantitativeValue

}

func (v CurrencyConversionService) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CurrencyConversionService"

	return json.Marshal(v)
}

func (v *CurrencyConversionService) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CurrencyConversionService"

	return json.Marshal(*v)
}