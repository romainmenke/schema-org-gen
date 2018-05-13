package schemaorggo

import "encoding/json"

// DepositAccount see : https://schema.org/DepositAccount
type DepositAccount struct {
	InvestmentOrDeposit

	typeContext

	// Amount see : https://schema.org/amount
	// The amount of money.
	// types : MonetaryAmount Number
	Amount []interface{} `json:"amount,omitempty"`
}

func (v DepositAccount) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.InvestmentOrDeposit.IntoMap(intop)

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

	*intop = into

	return nil
}

func (v DepositAccount) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DepositAccount"

	return data, nil
}

func (v DepositAccount) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
