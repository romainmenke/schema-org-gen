package schemaorg

import "encoding/json"

// PaymentCard see : https://schema.org/PaymentCard
type PaymentCard struct {

typeContext

FinancialProduct

// CashBack see : https://schema.orghttp://pending.schema.org/cashBack
// A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
CashBack interface{} `json:"cashBack"` // types : Boolean Number

// ContactlessPayment see : https://schema.orghttp://pending.schema.org/contactlessPayment
// A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
ContactlessPayment bool `json:"contactlessPayment"`

// FloorLimit see : https://schema.orghttp://pending.schema.org/floorLimit
// A floor limit is the amount of money above which credit card transactions must be authorized.
FloorLimit *MonetaryAmount `json:"floorLimit"`

}

func (v *PaymentCard) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentCard"

	return json.Marshal(v)
}
