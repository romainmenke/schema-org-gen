package schemaorggo

import "encoding/json"

// PaymentCard see : https://schema.org/PaymentCard
type PaymentCard struct {
	FinancialProduct

	typeContext

	// CashBack see : https://pending.schema.org/cashBack
	// A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
	// types : Boolean Number
	CashBack []interface{} `json:"cashBack,omitempty"`

	// ContactlessPayment see : https://pending.schema.org/contactlessPayment
	// A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
	// types : Boolean
	ContactlessPayment []bool `json:"contactlessPayment,omitempty"`

	// FloorLimit see : https://pending.schema.org/floorLimit
	// A floor limit is the amount of money above which credit card transactions must be authorized.
	// types : MonetaryAmount
	FloorLimit []*MonetaryAmount `json:"floorLimit,omitempty"`
}

func (v PaymentCard) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.FinancialProduct.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CashBack
		if len(v.CashBack) == 1 {
			value = v.CashBack[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cashBack"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContactlessPayment
		if len(v.ContactlessPayment) == 1 {
			value = v.ContactlessPayment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contactlessPayment"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FloorLimit
		if len(v.FloorLimit) == 1 {
			value = v.FloorLimit[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["floorLimit"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PaymentCard) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PaymentCard"

	return data, nil
}

func (v PaymentCard) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
