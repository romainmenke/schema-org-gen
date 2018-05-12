package schemaorggo

import "encoding/json"

// BankAccount see : https://schema.org/BankAccount
type BankAccount struct {
	FinancialProduct

	typeContext

	// AccountMinimumInflow see : http://pending.schema.org/accountMinimumInflow
	// A minimum amount that has to be paid in every month.
	// types : MonetaryAmount
	AccountMinimumInflow *MonetaryAmount `json:"accountMinimumInflow,omitempty"`

	// AccountOverdraftLimit see : http://pending.schema.org/accountOverdraftLimit
	// An overdraft is an extension of credit from a lending institution when an account reaches zero. An overdraft allows the individual to continue withdrawing money even if the account has no funds in it. Basically the bank allows people to borrow a set amount of money.
	// types : MonetaryAmount
	AccountOverdraftLimit *MonetaryAmount `json:"accountOverdraftLimit,omitempty"`

	// BankAccountType see : http://pending.schema.org/bankAccountType
	// The type of a bank account.
	// types : Text URL
	BankAccountType string `json:"bankAccountType,omitempty"`
}

func (v BankAccount) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BankAccount"

	return json.Marshal(v)
}

func (v *BankAccount) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BankAccount"

	return json.Marshal(*v)
}
