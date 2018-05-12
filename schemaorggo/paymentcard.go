package schemaorggo

import "encoding/json"

// PaymentCard see : https://schema.org/PaymentCard
type PaymentCard struct {
	FinancialProduct

	typeContext

	// CashBack see : http://pending.schema.org/cashBack
	// A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
	// types : Boolean Number
	CashBack interface{} `json:"cashBack,omitempty"`

	// ContactlessPayment see : http://pending.schema.org/contactlessPayment
	// A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
	// types : Boolean
	ContactlessPayment bool `json:"contactlessPayment,omitempty"`

	// FloorLimit see : http://pending.schema.org/floorLimit
	// A floor limit is the amount of money above which credit card transactions must be authorized.
	// types : MonetaryAmount
	FloorLimit *MonetaryAmount `json:"floorLimit,omitempty"`
}

func (v PaymentCard) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentCard"

	return json.Marshal(v)
}

func (v *PaymentCard) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PaymentCard"

	return json.Marshal(*v)
}
