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
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

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

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IsGift see : https://schema.org/isGift
	// Was the offer accepted as a gift for someone other than the buyer.
	// types : Boolean
	IsGift []bool `json:"isGift,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
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
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
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
