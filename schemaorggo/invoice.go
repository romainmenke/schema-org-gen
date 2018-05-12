package schemaorggo

import "encoding/json"

// Invoice see : https://schema.org/Invoice
type Invoice struct {
	Intangible

	typeContext

	// AccountId see : https://schema.org/accountId
	// The identifier for the account the payment will be applied to.
	// types : Text
	AccountId string `json:"accountId,omitempty"`

	// BillingPeriod see : https://schema.org/billingPeriod
	// The time interval used to compute the invoice.
	// types : Duration
	BillingPeriod *Duration `json:"billingPeriod,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	// types : Organization Person
	Broker interface{} `json:"broker,omitempty"`

	// Category see : https://schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : PhysicalActivityCategory Text Thing
	Category interface{} `json:"category,omitempty"`

	// ConfirmationNumber see : https://schema.org/confirmationNumber
	// A number that confirms the given order or payment has been received.
	// types : Text
	ConfirmationNumber string `json:"confirmationNumber,omitempty"`

	// Customer see : https://schema.org/customer
	// Party placing the order or paying the invoice.
	// types : Organization Person
	Customer interface{} `json:"customer,omitempty"`

	// MinimumPaymentDue see : https://schema.org/minimumPaymentDue
	// The minimum payment required at this time.
	// types : MonetaryAmount PriceSpecification
	MinimumPaymentDue interface{} `json:"minimumPaymentDue,omitempty"`

	// PaymentDueDate see : https://schema.org/paymentDueDate
	// The date that payment is due. Supersedes paymentDue (see: https://schema.org/paymentDue).
	// types : DateTime
	PaymentDueDate DateTime `json:"paymentDueDate,omitempty"`

	// PaymentMethod see : https://schema.org/paymentMethod
	// The name of the credit card or other method of payment for the order.
	// types : PaymentMethod
	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	// PaymentMethodId see : https://schema.org/paymentMethodId
	// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
	// types : Text
	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	// PaymentStatus see : https://schema.org/paymentStatus
	// The status of payment; whether the invoice has been paid or not.
	// types : PaymentStatusType Text
	PaymentStatus interface{} `json:"paymentStatus,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider interface{} `json:"provider,omitempty"`

	// ReferencesOrder see : https://schema.org/referencesOrder
	// The Order(s) related to this Invoice. One or more Orders may be combined into a single Invoice.
	// types : Order
	ReferencesOrder *Order `json:"referencesOrder,omitempty"`

	// ScheduledPaymentDate see : https://schema.org/scheduledPaymentDate
	// The date the invoice is scheduled to be paid.
	// types : Date
	ScheduledPaymentDate Date `json:"scheduledPaymentDate,omitempty"`

	// TotalPaymentDue see : https://schema.org/totalPaymentDue
	// The total amount due.
	// types : MonetaryAmount PriceSpecification
	TotalPaymentDue interface{} `json:"totalPaymentDue,omitempty"`
}

func (v Invoice) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Invoice"

	return json.Marshal(v)
}

func (v *Invoice) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Invoice"

	return json.Marshal(*v)
}
