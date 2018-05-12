package schemaorg

import "encoding/json"

// BankAccount see : https://schema.org/BankAccount
type BankAccount struct {

FinancialProduct

typeContext

// AccountMinimumInflow see : http://pending.schema.org/accountMinimumInflow
// A minimum amount that has to be paid in every month.
AccountMinimumInflow *MonetaryAmount `json:"accountMinimumInflow"`

// AccountOverdraftLimit see : http://pending.schema.org/accountOverdraftLimit
// An overdraft is an extension of credit from a lending institution when an account reaches zero. An overdraft allows the individual to continue withdrawing money even if the account has no funds in it. Basically the bank allows people to borrow a set amount of money.
AccountOverdraftLimit *MonetaryAmount `json:"accountOverdraftLimit"`

// BankAccountType see : http://pending.schema.org/bankAccountType
// The type of a bank account.
BankAccountType interface{} `json:"bankAccountType"` // types : Text URL

}

func (v *BankAccount) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BankAccount"

	return json.Marshal(v)
}
