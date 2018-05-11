package schemaorg

import "encoding/json"

// LoanOrCredit see : https://schema.org/LoanOrCredit
type LoanOrCredit struct {
	typeContext

	FinancialProduct

	// Amount see : /amount
	// The amount of money.
	Amount interface{} `json:"amount"` // types : MonetaryAmount Number

	// Currency see : /currency
	// The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
	Currency string `json:"currency"`

	// GracePeriod see : http://pending.schema.org/gracePeriod
	// The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
	GracePeriod *Duration `json:"gracePeriod"`

	// LoanRepaymentForm see : http://pending.schema.org/loanRepaymentForm
	// A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
	LoanRepaymentForm interface{} `json:"loanRepaymentForm"`

	// LoanTerm see : /loanTerm
	// The duration of the loan or credit agreement.
	LoanTerm *QuantitativeValue `json:"loanTerm"`

	// LoanType see : http://pending.schema.org/loanType
	// The type of a loan or credit.
	LoanType interface{} `json:"loanType"` // types : Text URL

	// RecourseLoan see : http://pending.schema.org/recourseLoan
	// The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
	RecourseLoan bool `json:"recourseLoan"`

	// RenegotiableLoan see : http://pending.schema.org/renegotiableLoan
	// Whether the terms for payment of interest can be renegotiated during the life of the loan.
	RenegotiableLoan bool `json:"renegotiableLoan"`

	// RequiredCollateral see : /requiredCollateral
	// Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
	RequiredCollateral interface{} `json:"requiredCollateral"` // types : Text Thing

}

func (v *LoanOrCredit) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LoanOrCredit"

	return json.Marshal(v)
}
