package schemaorggo

import "encoding/json"

// Order see : https://schema.org/Order
type Order struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AcceptedOffer see : https://schema.org/acceptedOffer
	// The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
	// types : Offer
	AcceptedOffer []*Offer `json:"acceptedOffer,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// BillingAddress see : https://schema.org/billingAddress
	// The billing address for the order.
	// types : PostalAddress
	BillingAddress []*PostalAddress `json:"billingAddress,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	// types : Person Organization
	Broker []interface{} `json:"broker,omitempty"`

	// ConfirmationNumber see : https://schema.org/confirmationNumber
	// A number that confirms the given order or payment has been received.
	// types : Text
	ConfirmationNumber []string `json:"confirmationNumber,omitempty"`

	// Customer see : https://schema.org/customer
	// Party placing the order or paying the invoice.
	// types : Organization Person
	Customer []interface{} `json:"customer,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Discount see : https://schema.org/discount
	// Any discount applied (to an Order).
	// types : Number Text
	Discount []interface{} `json:"discount,omitempty"`

	// DiscountCode see : https://schema.org/discountCode
	// Code used to redeem a discount.
	// types : Text
	DiscountCode []string `json:"discountCode,omitempty"`

	// DiscountCurrency see : https://schema.org/discountCurrency
	// The currency (in 3-letter ISO 4217 format) of the discount.
	// types : Text
	DiscountCurrency []string `json:"discountCurrency,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IsGift see : https://schema.org/isGift
	// Was the offer accepted as a gift for someone other than the buyer.
	// types : Boolean
	IsGift []bool `json:"isGift,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Merchant see : https://schema.org/merchant
	// &#39;merchant&#39; is an out-dated term for &#39;seller&#39;.
	// types : Organization Person
	Merchant []interface{} `json:"merchant,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// OrderDate see : https://schema.org/orderDate
	// Date order was placed.
	// types : DateTime
	OrderDate []DateTime `json:"orderDate,omitempty"`

	// OrderDelivery see : https://schema.org/orderDelivery
	// The delivery of the parcel related to this order or order item.
	// types : ParcelDelivery
	OrderDelivery []*ParcelDelivery `json:"orderDelivery,omitempty"`

	// OrderNumber see : https://schema.org/orderNumber
	// The identifier of the transaction.
	// types : Text
	OrderNumber []string `json:"orderNumber,omitempty"`

	// OrderStatus see : https://schema.org/orderStatus
	// The current status of the order.
	// types : OrderStatus
	OrderStatus []*OrderStatus `json:"orderStatus,omitempty"`

	// OrderedItem see : https://schema.org/orderedItem
	// The item ordered.
	// types : Product OrderItem
	OrderedItem []interface{} `json:"orderedItem,omitempty"`

	// PartOfInvoice see : https://schema.org/partOfInvoice
	// The order is being paid as part of the referenced Invoice.
	// types : Invoice
	PartOfInvoice []*Invoice `json:"partOfInvoice,omitempty"`

	// PaymentDue see : https://schema.org/paymentDue
	// The date that payment is due.
	// types : DateTime
	PaymentDue []DateTime `json:"paymentDue,omitempty"`

	// PaymentDueDate see : https://schema.org/paymentDueDate
	// The date that payment is due.
	// types : DateTime
	PaymentDueDate []DateTime `json:"paymentDueDate,omitempty"`

	// PaymentMethod see : https://schema.org/paymentMethod
	// The name of the credit card or other method of payment for the order.
	// types : PaymentMethod
	PaymentMethod []*PaymentMethod `json:"paymentMethod,omitempty"`

	// PaymentMethodId see : https://schema.org/paymentMethodId
	// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
	// types : Text
	PaymentMethodId []string `json:"paymentMethodId,omitempty"`

	// PaymentUrl see : https://schema.org/paymentUrl
	// The URL for sending a payment.
	// types : URL
	PaymentUrl []string `json:"paymentUrl,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	// types : Organization Person
	Seller []interface{} `json:"seller,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v Order) MarshalJSON() ([]byte, error) {
	type Alias Order

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Order\","), b[1:]...), nil
}
