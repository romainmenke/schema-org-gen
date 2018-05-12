package schemaorggo

import "encoding/json"

// Order see : https://schema.org/Order
type Order struct {
	Intangible

	typeContext

	// AcceptedOffer see : https://schema.org/acceptedOffer
	// The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
	// types : Offer
	AcceptedOffer *Offer `json:"acceptedOffer,omitempty"`

	// BillingAddress see : https://schema.org/billingAddress
	// The billing address for the order.
	// types : PostalAddress
	BillingAddress *PostalAddress `json:"billingAddress,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	// types : Organization Person
	Broker interface{} `json:"broker,omitempty"`

	// ConfirmationNumber see : https://schema.org/confirmationNumber
	// A number that confirms the given order or payment has been received.
	// types : Text
	ConfirmationNumber string `json:"confirmationNumber,omitempty"`

	// Customer see : https://schema.org/customer
	// Party placing the order or paying the invoice.
	// types : Organization Person
	Customer interface{} `json:"customer,omitempty"`

	// Discount see : https://schema.org/discount
	// Any discount applied (to an Order).
	// types : Number Text
	Discount interface{} `json:"discount,omitempty"`

	// DiscountCode see : https://schema.org/discountCode
	// Code used to redeem a discount.
	// types : Text
	DiscountCode string `json:"discountCode,omitempty"`

	// DiscountCurrency see : https://schema.org/discountCurrency
	// The currency (in 3-letter ISO 4217 format) of the discount.
	// types : Text
	DiscountCurrency string `json:"discountCurrency,omitempty"`

	// IsGift see : https://schema.org/isGift
	// Was the offer accepted as a gift for someone other than the buyer.
	// types : Boolean
	IsGift bool `json:"isGift,omitempty"`

	// OrderDate see : https://schema.org/orderDate
	// Date order was placed.
	// types : DateTime
	OrderDate DateTime `json:"orderDate,omitempty"`

	// OrderDelivery see : https://schema.org/orderDelivery
	// The delivery of the parcel related to this order or order item.
	// types : ParcelDelivery
	OrderDelivery *ParcelDelivery `json:"orderDelivery,omitempty"`

	// OrderNumber see : https://schema.org/orderNumber
	// The identifier of the transaction.
	// types : Text
	OrderNumber string `json:"orderNumber,omitempty"`

	// OrderStatus see : https://schema.org/orderStatus
	// The current status of the order.
	// types : OrderStatus
	OrderStatus *OrderStatus `json:"orderStatus,omitempty"`

	// OrderedItem see : https://schema.org/orderedItem
	// The item ordered.
	// types : OrderItem Product
	OrderedItem interface{} `json:"orderedItem,omitempty"`

	// PartOfInvoice see : https://schema.org/partOfInvoice
	// The order is being paid as part of the referenced Invoice.
	// types : Invoice
	PartOfInvoice *Invoice `json:"partOfInvoice,omitempty"`

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

	// PaymentUrl see : https://schema.org/paymentUrl
	// The URL for sending a payment.
	// types : URL
	PaymentUrl string `json:"paymentUrl,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller interface{} `json:"seller,omitempty"`
}

func (v Order) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Order"

	return json.Marshal(v)
}

func (v *Order) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Order"

	return json.Marshal(*v)
}
