package schemaorg

import "encoding/json"

// Invoice see : https://schema.org/Invoice
type Invoice struct {

typeContext

Intangible

// AccountId see : /accountId
// The identifier for the account the payment will be applied to.
AccountId string `json:"accountId"`

// BillingPeriod see : /billingPeriod
// The time interval used to compute the invoice.
BillingPeriod *Duration `json:"billingPeriod"`

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Organization Person

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
Category interface{} `json:"category"` // types : PhysicalActivityCategory Text Thing

// ConfirmationNumber see : /confirmationNumber
// A number that confirms the given order or payment has been received.
ConfirmationNumber string `json:"confirmationNumber"`

// Customer see : /customer
// Party placing the order or paying the invoice.
Customer interface{} `json:"customer"` // types : Organization Person

// MinimumPaymentDue see : /minimumPaymentDue
// The minimum payment required at this time.
MinimumPaymentDue interface{} `json:"minimumPaymentDue"` // types : MonetaryAmount PriceSpecification

// PaymentDueDate see : /paymentDueDate
// The date that payment is due. Supersedes paymentDue (see: https://schema.org/paymentDue).
PaymentDueDate interface{} `json:"paymentDueDate"`

// PaymentMethod see : /paymentMethod
// The name of the credit card or other method of payment for the order.
PaymentMethod *PaymentMethod `json:"paymentMethod"`

// PaymentMethodId see : /paymentMethodId
// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
PaymentMethodId string `json:"paymentMethodId"`

// PaymentStatus see : /paymentStatus
// The status of payment; whether the invoice has been paid or not.
PaymentStatus interface{} `json:"paymentStatus"` // types : PaymentStatusType Text

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// ReferencesOrder see : /referencesOrder
// The Order(s) related to this Invoice. One or more Orders may be combined into a single Invoice.
ReferencesOrder *Order `json:"referencesOrder"`

// ScheduledPaymentDate see : /scheduledPaymentDate
// The date the invoice is scheduled to be paid.
ScheduledPaymentDate interface{} `json:"scheduledPaymentDate"`

// TotalPaymentDue see : /totalPaymentDue
// The total amount due.
TotalPaymentDue interface{} `json:"totalPaymentDue"` // types : MonetaryAmount PriceSpecification

}

func (v *Invoice) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Invoice"

	return json.Marshal(v)
}
