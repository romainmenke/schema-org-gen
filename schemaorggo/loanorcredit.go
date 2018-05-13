package schemaorggo

import "encoding/json"

// LoanOrCredit see : https://schema.org/LoanOrCredit
type LoanOrCredit struct {
	FinancialProduct

	typeContext

	// Amount see : https://schema.org/amount
	// The amount of money.
	// types : MonetaryAmount Number
	Amount []interface{} `json:"amount,omitempty"`

	// Currency see : https://schema.org/currency
	// The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
	// types : Text
	Currency []string `json:"currency,omitempty"`

	// GracePeriod see : http://pending.schema.org/gracePeriod
	// The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
	// types : Duration
	GracePeriod []*Duration `json:"gracePeriod,omitempty"`

	// LoanRepaymentForm see : http://pending.schema.org/loanRepaymentForm
	// A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
	// types : RepaymentSpecification
	LoanRepaymentForm []interface{} `json:"loanRepaymentForm,omitempty"`

	// LoanTerm see : https://schema.org/loanTerm
	// The duration of the loan or credit agreement.
	// types : QuantitativeValue
	LoanTerm []*QuantitativeValue `json:"loanTerm,omitempty"`

	// LoanType see : http://pending.schema.org/loanType
	// The type of a loan or credit.
	// types : Text URL
	LoanType []string `json:"loanType,omitempty"`

	// RecourseLoan see : http://pending.schema.org/recourseLoan
	// The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
	// types : Boolean
	RecourseLoan []bool `json:"recourseLoan,omitempty"`

	// RenegotiableLoan see : http://pending.schema.org/renegotiableLoan
	// Whether the terms for payment of interest can be renegotiated during the life of the loan.
	// types : Boolean
	RenegotiableLoan []bool `json:"renegotiableLoan,omitempty"`

	// RequiredCollateral see : https://schema.org/requiredCollateral
	// Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
	// types : Text Thing
	RequiredCollateral []interface{} `json:"requiredCollateral,omitempty"`
}

func (v LoanOrCredit) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.FinancialProduct.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Amount
		if len(v.Amount) == 1 {
			value = v.Amount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["amount"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Currency
		if len(v.Currency) == 1 {
			value = v.Currency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["currency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GracePeriod
		if len(v.GracePeriod) == 1 {
			value = v.GracePeriod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gracePeriod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LoanRepaymentForm
		if len(v.LoanRepaymentForm) == 1 {
			value = v.LoanRepaymentForm[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["loanRepaymentForm"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LoanTerm
		if len(v.LoanTerm) == 1 {
			value = v.LoanTerm[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["loanTerm"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LoanType
		if len(v.LoanType) == 1 {
			value = v.LoanType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["loanType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecourseLoan
		if len(v.RecourseLoan) == 1 {
			value = v.RecourseLoan[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recourseLoan"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RenegotiableLoan
		if len(v.RenegotiableLoan) == 1 {
			value = v.RenegotiableLoan[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["renegotiableLoan"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredCollateral
		if len(v.RequiredCollateral) == 1 {
			value = v.RequiredCollateral[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredCollateral"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v LoanOrCredit) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "LoanOrCredit"

	return data, nil
}

func (v LoanOrCredit) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
