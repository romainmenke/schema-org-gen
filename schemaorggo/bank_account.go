package schemaorggo

import "encoding/json"

// BankAccount see : https://schema.org/BankAccount
type BankAccount struct {
	FinancialProduct

	typeContext

	// AccountMinimumInflow see : http://pending.schema.org/accountMinimumInflow
	// A minimum amount that has to be paid in every month.
	// types : MonetaryAmount
	AccountMinimumInflow []*MonetaryAmount `json:"accountMinimumInflow,omitempty"`

	// AccountOverdraftLimit see : http://pending.schema.org/accountOverdraftLimit
	// An overdraft is an extension of credit from a lending institution when an account reaches zero. An overdraft allows the individual to continue withdrawing money even if the account has no funds in it. Basically the bank allows people to borrow a set amount of money.
	// types : MonetaryAmount
	AccountOverdraftLimit []*MonetaryAmount `json:"accountOverdraftLimit,omitempty"`

	// BankAccountType see : http://pending.schema.org/bankAccountType
	// The type of a bank account.
	// types : Text URL
	BankAccountType []string `json:"bankAccountType,omitempty"`
}

func (v BankAccount) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.FinancialProduct.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AccountMinimumInflow
		if len(v.AccountMinimumInflow) == 1 {
			value = v.AccountMinimumInflow[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accountMinimumInflow"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccountOverdraftLimit
		if len(v.AccountOverdraftLimit) == 1 {
			value = v.AccountOverdraftLimit[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accountOverdraftLimit"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BankAccountType
		if len(v.BankAccountType) == 1 {
			value = v.BankAccountType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bankAccountType"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BankAccount) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BankAccount"

	return data, nil
}

func (v BankAccount) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
