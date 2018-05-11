package schemaorg

import "encoding/json"

// PaymentService see : https://schema.org/PaymentService
type PaymentService struct {

typeContext

FinancialProduct

// AnnualPercentageRate see : /annualPercentageRate
// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
AnnualPercentageRate interface{} `json:"annualPercentageRate"` // types : Number QuantitativeValue

// FeesAndCommissionsSpecification see : /feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

// InterestRate see : /interestRate
// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
InterestRate interface{} `json:"interestRate"` // types : Number QuantitativeValue

}

func (v *PaymentService) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentService"

	return json.Marshal(v)
}
