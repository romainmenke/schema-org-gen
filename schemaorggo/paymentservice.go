package schemaorggo

import "encoding/json"

// PaymentService see : https://schema.org/PaymentService
type PaymentService struct {
	FinancialProduct

	typeContext

	// AnnualPercentageRate see : https://schema.org/annualPercentageRate
	// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
	// types : Number QuantitativeValue
	AnnualPercentageRate interface{} `json:"annualPercentageRate,omitempty"`

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	// types : Text URL
	FeesAndCommissionsSpecification string `json:"feesAndCommissionsSpecification,omitempty"`

	// InterestRate see : https://schema.org/interestRate
	// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
	// types : Number QuantitativeValue
	InterestRate interface{} `json:"interestRate,omitempty"`
}

func (v PaymentService) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentService"

	return json.Marshal(v)
}

func (v *PaymentService) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PaymentService"

	return json.Marshal(*v)
}
