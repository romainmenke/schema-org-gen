package schemaorggo

import "encoding/json"

// LoanOrCredit see : https://schema.org/LoanOrCredit
type LoanOrCredit struct {
	FinancialProduct

	typeContext

	// Amount see : https://schema.org/amount
	// The amount of money.
	// types : MonetaryAmount Number
	Amount interface{} `json:"amount,omitempty"`

	// Currency see : https://schema.org/currency
	// The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
	// types : Text
	Currency string `json:"currency,omitempty"`

	// GracePeriod see : http://pending.schema.org/gracePeriod
	// The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
	// types : Duration
	GracePeriod *Duration `json:"gracePeriod,omitempty"`

	// LoanRepaymentForm see : http://pending.schema.org/loanRepaymentForm
	// A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
	// types : RepaymentSpecification
	LoanRepaymentForm interface{} `json:"loanRepaymentForm,omitempty"`

	// LoanTerm see : https://schema.org/loanTerm
	// The duration of the loan or credit agreement.
	// types : QuantitativeValue
	LoanTerm *QuantitativeValue `json:"loanTerm,omitempty"`

	// LoanType see : http://pending.schema.org/loanType
	// The type of a loan or credit.
	// types : Text URL
	LoanType string `json:"loanType,omitempty"`

	// RecourseLoan see : http://pending.schema.org/recourseLoan
	// The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
	// types : Boolean
	RecourseLoan bool `json:"recourseLoan,omitempty"`

	// RenegotiableLoan see : http://pending.schema.org/renegotiableLoan
	// Whether the terms for payment of interest can be renegotiated during the life of the loan.
	// types : Boolean
	RenegotiableLoan bool `json:"renegotiableLoan,omitempty"`

	// RequiredCollateral see : https://schema.org/requiredCollateral
	// Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
	// types : Text Thing
	RequiredCollateral interface{} `json:"requiredCollateral,omitempty"`
}

func (v LoanOrCredit) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LoanOrCredit"

	return json.Marshal(v)
}

func (v *LoanOrCredit) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LoanOrCredit"

	return json.Marshal(*v)
}
