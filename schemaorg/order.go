package schemaorg

import "encoding/json"

// Order see : https://schema.org/Order
type Order struct {

typeContext

Intangible

// AcceptedOffer see : https://schema.org/acceptedOffer
// The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
AcceptedOffer *Offer `json:"acceptedOffer"`

// BillingAddress see : https://schema.org/billingAddress
// The billing address for the order.
BillingAddress *PostalAddress `json:"billingAddress"`

// Broker see : https://schema.org/broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Organization Person

// ConfirmationNumber see : https://schema.org/confirmationNumber
// A number that confirms the given order or payment has been received.
ConfirmationNumber string `json:"confirmationNumber"`

// Customer see : https://schema.org/customer
// Party placing the order or paying the invoice.
Customer interface{} `json:"customer"` // types : Organization Person

// Discount see : https://schema.org/discount
// Any discount applied (to an Order).
Discount interface{} `json:"discount"` // types : Number Text

// DiscountCode see : https://schema.org/discountCode
// Code used to redeem a discount.
DiscountCode string `json:"discountCode"`

// DiscountCurrency see : https://schema.org/discountCurrency
// The currency (in 3-letter ISO 4217 format) of the discount.
DiscountCurrency string `json:"discountCurrency"`

// IsGift see : https://schema.org/isGift
// Was the offer accepted as a gift for someone other than the buyer.
IsGift bool `json:"isGift"`

// OrderDate see : https://schema.org/orderDate
// Date order was placed.
OrderDate interface{} `json:"orderDate"`

// OrderDelivery see : https://schema.org/orderDelivery
// The delivery of the parcel related to this order or order item.
OrderDelivery *ParcelDelivery `json:"orderDelivery"`

// OrderNumber see : https://schema.org/orderNumber
// The identifier of the transaction.
OrderNumber string `json:"orderNumber"`

// OrderStatus see : https://schema.org/orderStatus
// The current status of the order.
OrderStatus *OrderStatus `json:"orderStatus"`

// OrderedItem see : https://schema.org/orderedItem
// The item ordered.
OrderedItem interface{} `json:"orderedItem"` // types : OrderItem Product

// PartOfInvoice see : https://schema.org/partOfInvoice
// The order is being paid as part of the referenced Invoice.
PartOfInvoice *Invoice `json:"partOfInvoice"`

// PaymentDueDate see : https://schema.org/paymentDueDate
// The date that payment is due. Supersedes paymentDue (see: https://schema.org/paymentDue).
PaymentDueDate interface{} `json:"paymentDueDate"`

// PaymentMethod see : https://schema.org/paymentMethod
// The name of the credit card or other method of payment for the order.
PaymentMethod *PaymentMethod `json:"paymentMethod"`

// PaymentMethodId see : https://schema.org/paymentMethodId
// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
PaymentMethodId string `json:"paymentMethodId"`

// PaymentUrl see : https://schema.org/paymentUrl
// The URL for sending a payment.
PaymentUrl string `json:"paymentUrl"`

// Seller see : https://schema.org/seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
Seller interface{} `json:"seller"` // types : Organization Person

}

func (v *Order) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Order"

	return json.Marshal(v)
}
