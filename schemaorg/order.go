package schemaorg

// Order see : https://schema.org/Order
type Order struct {

Intangible// AcceptedOffer see : /acceptedOffer
// The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
AcceptedOffer string `json:"acceptedOffer"`

// BillingAddress see : /billingAddress
// The billing address for the order.
BillingAddress string `json:"billingAddress"`

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Organization Person

// ConfirmationNumber see : /confirmationNumber
// A number that confirms the given order or payment has been received.
ConfirmationNumber string `json:"confirmationNumber"`

// Customer see : /customer
// Party placing the order or paying the invoice.
Customer interface{} `json:"customer"` // types : Organization Person

// Discount see : /discount
// Any discount applied (to an Order).
Discount interface{} `json:"discount"` // types : Number Text

// DiscountCode see : /discountCode
// Code used to redeem a discount.
DiscountCode string `json:"discountCode"`

// DiscountCurrency see : /discountCurrency
// The currency (in 3-letter ISO 4217 format) of the discount.
DiscountCurrency string `json:"discountCurrency"`

// IsGift see : /isGift
// Was the offer accepted as a gift for someone other than the buyer.
IsGift string `json:"isGift"`

// OrderDate see : /orderDate
// Date order was placed.
OrderDate string `json:"orderDate"`

// OrderDelivery see : /orderDelivery
// The delivery of the parcel related to this order or order item.
OrderDelivery string `json:"orderDelivery"`

// OrderNumber see : /orderNumber
// The identifier of the transaction.
OrderNumber string `json:"orderNumber"`

// OrderStatus see : /orderStatus
// The current status of the order.
OrderStatus string `json:"orderStatus"`

// OrderedItem see : /orderedItem
// The item ordered.
OrderedItem interface{} `json:"orderedItem"` // types : OrderItem Product

// PartOfInvoice see : /partOfInvoice
// The order is being paid as part of the referenced Invoice.
PartOfInvoice string `json:"partOfInvoice"`

// PaymentDueDate see : /paymentDueDate
// The date that payment is due. Supersedes paymentDue (see: https://schema.org/paymentDue).
PaymentDueDate string `json:"paymentDueDate"`

// PaymentMethod see : /paymentMethod
// The name of the credit card or other method of payment for the order.
PaymentMethod string `json:"paymentMethod"`

// PaymentMethodId see : /paymentMethodId
// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
PaymentMethodId string `json:"paymentMethodId"`

// PaymentUrl see : /paymentUrl
// The URL for sending a payment.
PaymentUrl string `json:"paymentUrl"`

// Seller see : /seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
Seller interface{} `json:"seller"` // types : Organization Person

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// PartOfOrder see : /partOfOrder
// The overall order the items in this delivery were included in. 
PartOfOrder string `json:"partOfOrder"`

// ReferencesOrder see : /referencesOrder
// The Order(s) related to this Invoice. One or more Orders may be combined into a single Invoice. 
ReferencesOrder string `json:"referencesOrder"`

}

