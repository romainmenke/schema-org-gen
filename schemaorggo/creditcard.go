package schemaorggo

import "encoding/json"

// CreditCard see : https://schema.org/CreditCard
type CreditCard struct {
	LoanOrCredit

	typeContext

	// MonthlyMinimumRepaymentAmount see : http://pending.schema.org/monthlyMinimumRepaymentAmount
	// The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
	// types : MonetaryAmount Number
	MonthlyMinimumRepaymentAmount []interface{} `json:"monthlyMinimumRepaymentAmount,omitempty"`
}

func (v CreditCard) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.LoanOrCredit.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.MonthlyMinimumRepaymentAmount
		if len(v.MonthlyMinimumRepaymentAmount) == 1 {
			value = v.MonthlyMinimumRepaymentAmount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["monthlyMinimumRepaymentAmount"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v CreditCard) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CreditCard"

	return data, nil
}

func (v CreditCard) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
