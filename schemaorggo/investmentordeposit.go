package schemaorggo

import "encoding/json"

// InvestmentOrDeposit see : https://schema.org/InvestmentOrDeposit
type InvestmentOrDeposit struct {
	FinancialProduct

	typeContext

	// Amount see : https://schema.org/amount
	// The amount of money.
	// types : MonetaryAmount Number
	Amount []interface{} `json:"amount,omitempty"`
}

func (v InvestmentOrDeposit) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.FinancialProduct.intoMap(intop)

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

func (v InvestmentOrDeposit) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "InvestmentOrDeposit"

	return data, nil
}

func (v InvestmentOrDeposit) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
